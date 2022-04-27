package backuphandler

import (
	"context"
	"log"

	"github.com/opdev/backup-handler/cmd/command"
	"github.com/opdev/backup-handler/cmd/models"
	restoreservice "github.com/opdev/backup-handler/gen/restore_service"
)

// Restore Service service example implementation.
// The example methods log the requests and return zero values.
type restoreServicesrvc struct {
	logger *log.Logger
}

// NewRestoreService returns the Restore Service service implementation.
func NewRestoreService(logger *log.Logger) restoreservice.Service {
	return &restoreServicesrvc{logger}
}

// New restore request
func (s *restoreServicesrvc) Create(ctx context.Context, p *restoreservice.Restore) (res *restoreservice.Restoreresult, err error) {
	res = &restoreservice.Restoreresult{
		CreatedAt:            utcTime(),
		ID:                   genUUID(),
		Name:                 p.Name,
		Namespace:            p.Namespace,
		StorageSecret:        p.StorageSecret,
		DestinationName:      p.DestinationName,
		DestinationNamespace: p.DestinationNamespace,
		BackupLocation:       p.BackupLocation,
	}
	s.logger.Print("restoreService.create")

	if err = models.CreateRestore(res); err != nil {
		return nil, err
	}

	if err := command.StartRestore(res); err != nil {
		return nil,
			&restoreservice.BackupNotFound{
				Message: "backup resource not found.",
			}
	}

	return
}

// Get restore request
func (s *restoreServicesrvc) Get(ctx context.Context, p *restoreservice.GetPayload) (res *restoreservice.Restoreresult, err error) {
	res = &restoreservice.Restoreresult{
		ID: p.ID,
	}
	s.logger.Print("restoreService.get")

	if err = models.GetRestore(res); err != nil {
		return nil, err
	}

	return
}

// Update restore request
func (s *restoreServicesrvc) Update(ctx context.Context, p *restoreservice.Restoreresult) (res *restoreservice.Restoreresult, err error) {
	p.UpdatedAt = utcTime()
	res = p
	s.logger.Print("restoreService.update")
	return
}

// Mark complete restore request
func (s *restoreServicesrvc) Delete(ctx context.Context, p *restoreservice.DeletePayload) (res *restoreservice.Restoreresult, err error) {
	res = &restoreservice.Restoreresult{
		ID:        p.ID,
		DeletedAt: utcTime(),
	}
	s.logger.Print("restoreService.delete")
	return
}
