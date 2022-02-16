package command

import (
	"encoding/json"
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

		log.Printf("%d backups received", len(backups.Items))
		for _, backup := range backups.Items {
			results, err := execBackup(backup)
			if err != nil {
				log.Fatalf("error running backup; %v\n", err)
			}

			log.Printf("stdout: %s | stderr: %s", results.Output(), results.Error())

		}
		time.Sleep(3 * time.Second)
	}
}

// func markBackupRunning() {}

// func markBackupCompleted() {}

// func isAlive() {
// 	backOffTimes := []time.Duration{
// 		1 * time.Second,
// 		3 * time.Second,
// 		5 * time.Second,
// 	}

// 	for _, backoff := range backOffTimes {
// 		func() {
// 			_, err := http.Get("http://localhost:8890/next-batch")
// 			if err != nil {
// 				time.Sleep(backoff)
// 			}
// 		}()
// 	}
// }

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

func execBackup(backup backupv1.Backup) (*ExecResponse, error) {
	return ExecuteCommand(
		ExecOptions{
			Pod:       backup.PodName,
			Container: backup.ContainerName,
			Namespace: backup.Namespace,
			Command:   backup.Command,
		},
	)
}
