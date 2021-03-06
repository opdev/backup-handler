// Code generated by goa v3.7.3, DO NOT EDIT.
//
// Restore Service service
//
// Command:
// $ goa gen github.com/opdev/backup-handler/design

package restoreservice

import (
	"context"

	restoreserviceviews "github.com/opdev/backup-handler/gen/restore_service/views"
)

// Service to handle restore requests
type Service interface {
	// New restore request
	Create(context.Context, *Restore) (res *Restoreresult, err error)
	// Get restore request
	Get(context.Context, *GetPayload) (res *Restoreresult, err error)
	// Update restore request
	Update(context.Context, *Restoreresult) (res *Restoreresult, err error)
	// Mark complete restore request
	Delete(context.Context, *DeletePayload) (res *Restoreresult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Restore Service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"create", "get", "update", "delete"}

// Backup not found error is returned when the backup is not found.
type BackupNotFound struct {
	// backup resource not found
	Message string
}

// DeletePayload is the payload type of the Restore Service service delete
// method.
type DeletePayload struct {
	ID *string
}

// GetPayload is the payload type of the Restore Service service get method.
type GetPayload struct {
	ID *string
}

// Restore is the payload type of the Restore Service service create method.
type Restore struct {
	// Name of pachyderm instance to restore to
	Name *string
	// Namespace to restore to
	Namespace *string
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string
	// Key of backup tarball
	BackupLocation *string
}

// Restoreresult is the result type of the Restore Service service create
// method.
type Restoreresult struct {
	CreatedAt *string
	UpdatedAt *string
	DeletedAt *string
	ID        *string
	// Name of pachyderm instance to restore to
	Name *string
	// Namespace to restore to
	Namespace *string
	// Key of backup tarball
	BackupLocation *string
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string
	// base64 encoded kubernetes object
	KubernetesResource *string
	// base64 encoded database dump
	Database *string
}

// Error returns an error description.
func (e *BackupNotFound) Error() string {
	return "Backup not found error is returned when the backup is not found."
}

// ErrorName returns "BackupNotFound".
func (e *BackupNotFound) ErrorName() string {
	return "backup_not_found"
}

// NewRestoreresult initializes result type Restoreresult from viewed result
// type Restoreresult.
func NewRestoreresult(vres *restoreserviceviews.Restoreresult) *Restoreresult {
	return newRestoreresult(vres.Projected)
}

// NewViewedRestoreresult initializes viewed result type Restoreresult from
// result type Restoreresult using the given view.
func NewViewedRestoreresult(res *Restoreresult, view string) *restoreserviceviews.Restoreresult {
	p := newRestoreresultView(res)
	return &restoreserviceviews.Restoreresult{Projected: p, View: "default"}
}

// newRestoreresult converts projected type Restoreresult to service type
// Restoreresult.
func newRestoreresult(vres *restoreserviceviews.RestoreresultView) *Restoreresult {
	res := &Restoreresult{
		CreatedAt:          vres.CreatedAt,
		UpdatedAt:          vres.UpdatedAt,
		DeletedAt:          vres.DeletedAt,
		ID:                 vres.ID,
		Name:               vres.Name,
		Namespace:          vres.Namespace,
		BackupLocation:     vres.BackupLocation,
		StorageSecret:      vres.StorageSecret,
		KubernetesResource: vres.KubernetesResource,
		Database:           vres.Database,
	}
	return res
}

// newRestoreresultView projects result type Restoreresult to projected type
// RestoreresultView using the "default" view.
func newRestoreresultView(res *Restoreresult) *restoreserviceviews.RestoreresultView {
	vres := &restoreserviceviews.RestoreresultView{
		CreatedAt:          res.CreatedAt,
		UpdatedAt:          res.UpdatedAt,
		DeletedAt:          res.DeletedAt,
		ID:                 res.ID,
		Name:               res.Name,
		Namespace:          res.Namespace,
		BackupLocation:     res.BackupLocation,
		StorageSecret:      res.StorageSecret,
		KubernetesResource: res.KubernetesResource,
		Database:           res.Database,
	}
	return vres
}
