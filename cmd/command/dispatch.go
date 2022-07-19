package command

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	backupservice "github.com/opdev/backup-handler/gen/backup_service"
	backupclient "github.com/opdev/backup-handler/gen/http/backup_service/client"
	"github.com/walle/targz"
	"k8s.io/apimachinery/pkg/types"
)

// RunBackup executs a backup job
func RunBackup(backup *backupservice.Backupresult) error {
	logger := log.New(os.Stderr, "[backuprunner] ", log.Ldate|log.Ltime)
	logger.Println("Starting backup..")
	results, err := execBackup(backup)
	if err != nil {
		return err
	}

	if err := writeBackup(backup, results.Output()); err != nil {
		return err
	}

	if stderr := results.Error(); stderr != "" {
		logger.Printf("error executing backup; %s\n", stderr)
	}

	if err := setBackupResults(backup); err != nil {
		return err
	}

	return markBackupCompleted(logger, backup)
}

func writeBackup(backup *backupservice.Backupresult, output string) error {
	var t time.Time = time.Now().UTC()
	timestamp := t.Format("200601021504")
	backupDir := path.Join(
		os.TempDir(),
		strings.Join([]string{"pachyderm-backup", timestamp}, "-"),
	)
	backupTarball := path.Join(
		os.TempDir(),
		fmt.Sprintf("%s.tar.gz", backupDir),
	)

	// create temp directory to hold backup
	if err := os.Mkdir(backupDir, 0750); err != nil {
		return err
	}

	// write the custom resource to the file cr.yaml
	crData, err := base64.StdEncoding.DecodeString(*backup.KubernetesResource)
	if err != nil {
		return err
	}
	cr, err := os.Create(path.Join(backupDir, "cr.json"))
	if err != nil {
		return err
	}
	if _, err := cr.Write(crData); err != nil {
		return err
	}
	defer cr.Close()

	dump, err := os.Create(path.Join(backupDir, "database.tar"))
	if err != nil {
		return err
	}
	if _, err := dump.Write([]byte(output)); err != nil {
		return err
	}
	defer dump.Close()

	if err := targz.Compress(backupDir, backupTarball); err != nil {
		return err
	}

	response, err := uploadBackup(*backup.StorageSecret, *backup.Namespace, backupTarball)
	if err != nil {
		return err
	}
	backup.Location = &response.Location

	if err = os.Remove(backupTarball); err != nil {
		return err
	}

	if err := os.RemoveAll(backupDir); err != nil {
		return err
	}

	return nil
}

func setBackupResults(backup *backupservice.Backupresult) error {
	conversion := backupclient.UpdateResponseBody(*backup)
	payload, err := json.Marshal(&conversion)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, "http://localhost:8890/backups", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	defer request.Body.Close()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &backup); err != nil {
		return err
	}

	return nil
}

func markBackupCompleted(logger *log.Logger, backup *backupservice.Backupresult) error {
	url := fmt.Sprintf("http://localhost:8890/backups/%s", *backup.ID)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &backup); err != nil {
		return err
	}

	logger.Println("backup marked completed!")

	return nil
}

func execBackup(backup *backupservice.Backupresult) (*ExecResponse, error) {
	cmdBuilder := &Builder{
		cmd: *backup.Command,
	}
	cmd, err := cmdBuilder.Unmarshal()
	if err != nil {
		return nil, err
	}

	return ExecuteCommand(
		ExecOptions{
			Pod:       *backup.Pod,
			Container: *backup.Container,
			Namespace: *backup.Namespace,
			Command:   cmd,
		},
	)
}

func uploadBackup(name, namespace, backupName string) (*s3manager.UploadOutput, error) {
	creds, err := getUploadCredentials(
		types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
	)
	if err != nil {
		return nil, err
	}

	_session, err := creds.awsSession()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(backupName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	uploader := s3manager.NewUploader(_session)
	return uploader.Upload(
		&s3manager.UploadInput{
			Bucket: aws.String(creds.bucket),
			Key:    aws.String(path.Join("/", "backups", filepath.Base(f.Name()))),
			Body:   f,
		},
	)
}

func (r *uploadCredentials) awsSession() (*session.Session, error) {
	return session.NewSession(
		&aws.Config{
			Region: aws.String(r.region),
			Credentials: credentials.NewStaticCredentials(
				r.accessID,
				r.accessSecret,
				"",
			),
		},
	)
}
