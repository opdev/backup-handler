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
	backupv1 "github.com/opdev/backup-handler/api/v1"
	"github.com/walle/targz"
	"k8s.io/apimachinery/pkg/types"
)

func BackupDispatch() {
	time.Sleep(10 * time.Second)

	for {
		backups := fetchBackupBatch()

		for _, backup := range backups.Items {
			results, err := execBackup(backup)
			if err != nil {
				log.Fatalf("error running backup; %v\n", err)
			}

			if err := writeBackup(backup, results.Output()); err != nil {
				log.Printf("error writing backup.\n%v.\n", err)
			}

			if stderr := results.Error(); stderr != "" {
				log.Printf("stderr: %s\n", stderr)
			}

			if err := setBackupResults(backup); err != nil {
				log.Println("error updating backup response.", err.Error())
			}

			if err := markBackupCompleted(backup); err != nil {
				log.Fatal("error marking backup completed")
			}

			log.Printf("backup %s has completed.", backup.ID.String())
		}
		time.Sleep(3 * time.Second)
	}
}

func writeBackup(backup *backupv1.Backup, output string) error {
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
	crData, err := base64.StdEncoding.DecodeString(backup.Resource)
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

	dump, err := os.Create(path.Join(backupDir, "database.sql"))
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

	response, err := uploadBackup(backup.UploadSecret, backup.Namespace, backupTarball)
	if err != nil {
		return err
	}
	backup.UploadLocation = response.Location

	if err = os.Remove(backupTarball); err != nil {
		return err
	}

	if err := os.RemoveAll(backupDir); err != nil {
		return err
	}

	return nil
}

func setBackupResults(backup *backupv1.Backup) error {
	payload, err := json.Marshal(backup)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, "http://localhost:8890/backup", bytes.NewBuffer(payload))
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

func markBackupCompleted(backup *backupv1.Backup) error {
	url := fmt.Sprintf("http://localhost:8890/backup/%s", backup.ID.String())
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

	return nil
}

func fetchBackupBatch() backupv1.BackupList {
	resp, err := http.Get("http://localhost:8890/next-batch")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error retrieving backup jobs: %v\n", err)
	}

	backups := backupv1.BackupList{}
	if err := json.Unmarshal(body, &backups); err != nil {
		log.Fatalf("error unmarshaling json to struct. %v\n", err)
	}

	return backups
}

func execBackup(backup *backupv1.Backup) (*ExecResponse, error) {
	return ExecuteCommand(
		ExecOptions{
			Pod:       backup.PodName,
			Container: backup.ContainerName,
			Namespace: backup.Namespace,
			Command:   backup.Command,
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
