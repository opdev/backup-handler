// Code generated by goa v3.7.3, DO NOT EDIT.
//
// Backup Service service
//
// Command:
// $ goa gen github.com/opdev/backup-handler/design

package backupservice

import (
	"context"

	backupserviceviews "github.com/opdev/backup-handler/gen/backup_service/views"
)

// Service to handle backup requests
type Service interface {
	// New backup request
	Create(context.Context, *Backup) (res *Backupresult, err error)
	// Obtain backup request
	Get(context.Context, *GetPayload) (res *Backupresult, err error)
	// Update backup request
	Update(context.Context, *Backupresult) (res *Backupresult, err error)
	// Mark complete backup request
	Delete(context.Context, *DeletePayload) (res *Backupresult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Backup Service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"create", "get", "update", "delete"}

// Backup is the payload type of the Backup Service service create method.
type Backup struct {
	// Name of pachyderm instance backed up
	Name *string
	// Namespace of resource backed up
	Namespace *string
	// Name of target pod
	Pod *string
	// Name of container in pod
	Container *string
	// base64 encoded command to run in pod
	Command *string
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string
	// base64 encoded json representation of object
	KubernetesResource *string
}

// Backup not found error is returned when the backup is not found.
type BackupNotFound struct {
	// backup resource not found
	Message string
}

// Backupresult is the result type of the Backup Service service create method.
type Backupresult struct {
	CreatedAt *string
	UpdatedAt *string
	DeletedAt *string
	ID        *string
	// Current state of the job
	State *string
	// Name of pachyderm instance backed up
	Name *string
	// Namespace of resource backed up
	Namespace *string
	// Name of target pod
	Pod *string
	// Name of container in pod
	Container *string
	// base64 encoded command to run in pod
	Command *string
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string
	// base64 encoded json representation of object
	KubernetesResource *string
	// URL of the uploaded backup tarball
	Location *string
}

// DeletePayload is the payload type of the Backup Service service delete
// method.
type DeletePayload struct {
	ID *string
}

// GetPayload is the payload type of the Backup Service service get method.
type GetPayload struct {
	ID *string
}

// Error returns an error description.
func (e *BackupNotFound) Error() string {
	return "Backup not found error is returned when the backup is not found."
}

// ErrorName returns "BackupNotFound".
func (e *BackupNotFound) ErrorName() string {
	return "backup_not_found"
}

// NewBackupresult initializes result type Backupresult from viewed result type
// Backupresult.
func NewBackupresult(vres *backupserviceviews.Backupresult) *Backupresult {
	return newBackupresult(vres.Projected)
}

// NewViewedBackupresult initializes viewed result type Backupresult from
// result type Backupresult using the given view.
func NewViewedBackupresult(res *Backupresult, view string) *backupserviceviews.Backupresult {
	p := newBackupresultView(res)
	return &backupserviceviews.Backupresult{Projected: p, View: "default"}
}

// newBackupresult converts projected type Backupresult to service type
// Backupresult.
func newBackupresult(vres *backupserviceviews.BackupresultView) *Backupresult {
	res := &Backupresult{
		CreatedAt:          vres.CreatedAt,
		UpdatedAt:          vres.UpdatedAt,
		DeletedAt:          vres.DeletedAt,
		ID:                 vres.ID,
		State:              vres.State,
		Name:               vres.Name,
		Namespace:          vres.Namespace,
		Pod:                vres.Pod,
		Container:          vres.Container,
		Command:            vres.Command,
		StorageSecret:      vres.StorageSecret,
		KubernetesResource: vres.KubernetesResource,
		Location:           vres.Location,
	}
	return res
}

// newBackupresultView projects result type Backupresult to projected type
// BackupresultView using the "default" view.
func newBackupresultView(res *Backupresult) *backupserviceviews.BackupresultView {
	vres := &backupserviceviews.BackupresultView{
		CreatedAt:          res.CreatedAt,
		UpdatedAt:          res.UpdatedAt,
		DeletedAt:          res.DeletedAt,
		ID:                 res.ID,
		State:              res.State,
		Name:               res.Name,
		Namespace:          res.Namespace,
		Pod:                res.Pod,
		Container:          res.Container,
		Command:            res.Command,
		StorageSecret:      res.StorageSecret,
		KubernetesResource: res.KubernetesResource,
		Location:           res.Location,
	}
	return vres
}
