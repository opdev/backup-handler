package backuphandler

import (
	"context"
	"fmt"
	"log"

	"github.com/opdev/backup-handler/cmd/models"
	backupservice "github.com/opdev/backup-handler/gen/backup_service"
)

// Initialize backup queue
var queue *BackupQueue = &BackupQueue{}

// Backup Service service example implementation.
// The example methods log the requests and return zero values.
type backupServicesrvc struct {
	logger *log.Logger
}

// NewBackupService returns the Backup Service service implementation.
func NewBackupService(logger *log.Logger) backupservice.Service {
	return &backupServicesrvc{logger}
}

// New backup request
func (s *backupServicesrvc) Create(ctx context.Context, p *backupservice.Backup) (res *backupservice.Backupresult, err error) {
	res = newBackup(p)
	s.logger.Print("backupService.create")

	if err := models.CreateBackup(res); err != nil {
		return nil, err
	}

	queue.Add(res)
	go queue.execBackup(s.logger)

	return
}

// Obtain backup request
func (s *backupServicesrvc) Get(ctx context.Context, p *backupservice.GetPayload) (res *backupservice.Backupresult, err error) {
	res = &backupservice.Backupresult{
		ID: p.ID,
	}
	s.logger.Print("backupService.get")

	if err := models.GetBackup(res); err != nil {
		return nil, &backupservice.BackupNotFound{
			Message: fmt.Sprintf("backup ID %s not found.", *p.ID),
		}
	}

	return
}

// Update backup request
func (s *backupServicesrvc) Update(ctx context.Context, p *backupservice.Backupresult) (res *backupservice.Backupresult, err error) {
	res = &backupservice.Backupresult{
		UpdatedAt: utcTime(),
	}
	s.logger.Print("backupService.update")

	_, err = models.UpdateBackup(res)
	if err != nil {
		return nil, &backupservice.BackupNotFound{
			Message: fmt.Sprintf("backup ID %s not found.", *p.ID),
		}
	}

	return
}

// Mark complete backup request
func (s *backupServicesrvc) Delete(ctx context.Context, p *backupservice.DeletePayload) (res *backupservice.Backupresult, err error) {
	res = &backupservice.Backupresult{
		DeletedAt: utcTime(),
	}
	s.logger.Print("backupService.delete")

	_, err = models.DeleteBackup(res)
	if err != nil {
		return nil, &backupservice.BackupNotFound{
			Message: fmt.Sprintf("backup ID %s not found.", *p.ID),
		}
	}

	return
}

func newBackup(p *backupservice.Backup) *backupservice.Backupresult {
	state := "queued"
	return &backupservice.Backupresult{
		ID:                 genUUID(),
		CreatedAt:          utcTime(),
		Name:               p.Name,
		Namespace:          p.Namespace,
		Pod:                p.Pod,
		Container:          p.Container,
		Command:            p.Command,
		State:              &state,
		StorageSecret:      p.StorageSecret,
		KubernetesResource: p.KubernetesResource,
	}
}
