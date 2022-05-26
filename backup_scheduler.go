package backuphandler

import (
	"log"
	"sync"

	"github.com/opdev/backup-handler/cmd/command"
	"github.com/opdev/backup-handler/cmd/models"
	backupservice "github.com/opdev/backup-handler/gen/backup_service"
)

// BackupQueue holds a slice of backups to be ran
type BackupQueue struct {
	backups []*backupservice.Backupresult
	mu      sync.Mutex
}

// Add method adds a backup requests to the backup queue
func (r *BackupQueue) Add(backup *backupservice.Backupresult) {
	r.backups = append(r.backups, backup)
}

// Pop method returns the first item in the backup queue
func (r *BackupQueue) Pop() *backupservice.Backupresult {
	r.mu.Lock()
	defer r.mu.Unlock()
	resp := r.backups[0]
	r.backups = r.backups[1:]
	return resp
}

func (r *BackupQueue) execBackup(logger *log.Logger) {
	for len(r.backups) > 0 {
		target := r.Pop()
		state := "running"
		target.UpdatedAt = utcTime()
		target.State = &state
		// mark backup as running
		if _, err := models.UpdateBackup(target); err != nil {
			logger.Printf("error updating backup state")
		}

		defer logger.Printf("backup %s has completed.\n", *target.ID)
		// do the work
		if err := command.RunBackup(target); err != nil {
			logger.Printf("error running backup %s. %+v\n", *target.ID, err)
		}

		// mark the backup job completed
		state = "completed"
		target.DeletedAt = utcTime()
		target.State = &state
		if _, err := models.UpdateBackup(target); err != nil {
			logger.Printf("error updating backup state")
		}
	}
}
