package command

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/kubectl/pkg/scheme"
)

func kubeConfig() (*rest.Config, error) {
	conf, err := rest.InClusterConfig()
	if err != nil {
		cfg := filepath.Join("~/", ".kube", "config")
		if os.Getenv("KUBECONFIG") != "" {
			cfg = os.Getenv("KUBECONFIG")
		}
		return clientcmd.BuildConfigFromFlags("", cfg)
	}

	return conf, nil
}

func kubeClient(config *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(config)
}

type APIClient struct {
	config  *rest.Config
	clients *kubernetes.Clientset
}

// NewAPIClient allow users to get a way to interact with Openshift
func NewAPIClient() (*APIClient, error) {
	config, err := kubeConfig()
	if err != nil {
		return nil, err
	}

	clients, err := kubeClient(config)
	if err != nil {
		return nil, err
	}

	return &APIClient{
		config:  config,
		clients: clients,
	}, nil
}

func (r *APIClient) GetPod(ctx context.Context, pod types.NamespacedName) (*corev1.Pod, error) {
	return r.clients.CoreV1().Pods(pod.Namespace).Get(ctx, pod.Name, metav1.GetOptions{})
}

// Client returns the kubernetes clientsets
func (r *APIClient) Client() *kubernetes.Clientset {
	return r.clients
}

// Config returns the kubernetes rest config
func (r *APIClient) Config() *rest.Config {
	return r.config
}

// ExecResponse provides a mechanism for an
// executed command to return output following the
// completion of a command
type ExecResponse struct {
	stdout string
	stderr string
}

// ExecOptions provides options to be
// used when executing commands in poda
type ExecOptions struct {
	Pod       string
	Container string
	Namespace string
	Command   []string
}

// Output returns the command output in string format
func (r *ExecResponse) Output() string {
	return r.stdout
}

// Error method returns error in string format
func (r *ExecResponse) Error() string {
	return r.stderr
}

// ExecuteCommand takes a command and runs it on a specific pod
func ExecuteCommand(options ExecOptions) (*ExecResponse, error) {
	config, err := kubeConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubeClient(config)
	if err != nil {
		return nil, err
	}

	request := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(options.Pod).
		Namespace(options.Namespace).
		SubResource("exec").
		Param("container", options.Container)
	request.VersionedParams(
		&corev1.PodExecOptions{
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
			Container: options.Container,
			Command:   options.Command,
		}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", request.URL())
	if err != nil {
		return nil, err
	}

	var stdout, stderr bytes.Buffer
	err = exec.Stream(
		remotecommand.StreamOptions{
			Stdin:  nil,
			Stdout: &stdout,
			Stderr: &stderr,
		},
	)
	if err != nil {
		return nil, err
	}

	return &ExecResponse{
		stdout: stdout.String(),
		stderr: stderr.String(),
	}, nil
}

func getSecret(secret types.NamespacedName) (*corev1.Secret, error) {
	config, err := kubeConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubeClient(config)
	if err != nil {
		return nil, err
	}

	return client.
		CoreV1().
		Secrets(secret.Namespace).
		Get(context.Background(), secret.Name,
			metav1.GetOptions{
				TypeMeta: metav1.TypeMeta{
					Kind: "Secret",
				},
			},
		)
}

type uploadCredentials struct {
	bucket       string
	region       string
	accessID     string
	accessSecret string
}

func getUploadCredentials(secretKey types.NamespacedName) (*uploadCredentials, error) {
	secret, err := getSecret(secretKey)
	if err != nil {
		return nil, err
	}

	uploadCreds := &uploadCredentials{}
	bucket, ok := secret.Data["bucket"]
	if !ok {
		return nil, errors.New("missing S3 bucket name")
	}
	uploadCreds.bucket = string(bucket)

	region, ok := secret.Data["region"]
	if !ok {
		return nil, errors.New("missing S3 bucket")
	}
	uploadCreds.region = string(region)

	accessID, ok := secret.Data["access-id"]
	if !ok {
		return nil, errors.New("missing S3 access ID")
	}
	uploadCreds.accessID = string(accessID)

	accessSecret, ok := secret.Data["access-secret"]
	if !ok {
		return nil, errors.New("missing S3 access secret key")
	}
	uploadCreds.accessSecret = string(accessSecret)

	return uploadCreds, nil
}
