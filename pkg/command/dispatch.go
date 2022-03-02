package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	backupv1 "github.com/opdev/backup-handler/api/v1"
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

			fmt.Println("stdout: ", results.Output())
			fmt.Println("stderr: ", results.Error())
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
