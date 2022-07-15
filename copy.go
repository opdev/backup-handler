package backuphandler

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/opdev/backup-handler/cmd/command"
	restoreservice "github.com/opdev/backup-handler/gen/restore_service"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/cp"
	"k8s.io/kubectl/pkg/scheme"
)

func restoreDatabase(restore *restoreservice.Restoreresult) error {
	var (
		namespace string
		container string
		pod       string
	)
	{
		namespace = "pachyderm-testing"
		container = "postgres"
		pod = "postgres-0"
	}

	// write dump to file
	dbDump, err := func(db string) (string, error) {
		tmpdir, err := os.MkdirTemp(path.Join("/", "tmp"), "db-")
		if err != nil {
			return "", err
		}

		f, err := os.Create(path.Join(tmpdir, "dump.sql"))
		if err != nil {
			return "", err
		}
		defer f.Close()

		dump, err := base64.StdEncoding.DecodeString(db)
		if err != nil {
			return "", err
		}

		_, err = f.Write(dump)
		if err != nil {
			return "", err
		}

		return f.Name(), nil
	}(*restore.Database)
	if err != nil {
		return err
	}

	if err := Copy(pod, container, namespace, dbDump); err != nil {
		return err
	}

	response, err := func(pod, container, namespace string) (*command.ExecResponse, error) {
		res, err := command.ExecuteCommand(
			command.ExecOptions{
				Pod:       pod,
				Container: container,
				Namespace: namespace,
				Command:   []string{"psql", "pachyderm", "-f", "/tmp/restore.sql"},
			})
		if err != nil {
			return nil, err
		}

		return res, nil
	}(pod, container, namespace)
	if err != nil {
		return err
	}

	if response.Output() != "" {
		fmt.Println(response.Output())
	}

	if response.Error() != "" {
		fmt.Println(response.Error())
	}

	return nil
}

// Copy function allows the copying of files to and from pods.
// Similar to "kubectl cp" in implementation
// The function accepts, pod, container, namespace and name of file to be
// copied as arguments and returns and error
func Copy(pod, container, namespace, filename string) error {
	api, err := command.NewAPIClient()
	if err != nil {
		return err
	}

	config, client := api.Config(), api.Client()
	config.APIPath = "/api"
	config.GroupVersion = &schema.GroupVersion{Version: "v1"}
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}

	ioStream, stdin, stdout, stderr := genericclioptions.NewTestIOStreams()
	cpOptions := cp.NewCopyOptions(ioStream)
	cpOptions.ClientConfig = config
	cpOptions.Clientset = client
	cpOptions.Container = container
	cpOptions.Namespace = namespace

	destination := fmt.Sprintf("%s/%s:/tmp/%s", namespace, pod, path.Base(filename))
	if err := cpOptions.Run([]string{filename, destination}); err != nil {
		log.Fatal("error executing copy; ", err)
	}

	// perform database dump on backup-handler pod
	if err := os.RemoveAll(path.Dir(filename)); err != nil {
		return err
	}

	if stdin != nil {
		fmt.Println(stdin.String())
	}

	if stdout != nil {
		fmt.Println(stdout.String())
	}

	if stderr != nil {
		fmt.Println(stderr.String())
	}

	return nil
}
