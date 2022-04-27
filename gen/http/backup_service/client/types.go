// Code generated by goa v3.7.2, DO NOT EDIT.
//
// Backup Service HTTP client types
//
// Command:
// $ goa gen github.com/opdev/backup-handler/design

package client

import (
	backupservice "github.com/opdev/backup-handler/gen/backup_service"
	backupserviceviews "github.com/opdev/backup-handler/gen/backup_service/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "Backup Service" service "create"
// endpoint HTTP request body.
type CreateRequestBody struct {
	// Name of pachyderm instance backed up
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Namespace of resource backed up
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Name of target pod
	Pod *string `form:"pod,omitempty" json:"pod,omitempty" xml:"pod,omitempty"`
	// Name of container in pod
	Container *string `form:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// base64 encoded command to run in pod
	Command *string `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string `form:"storage_secret,omitempty" json:"storage_secret,omitempty" xml:"storage_secret,omitempty"`
	// base64 encoded json representation of object
	KubernetesResource *string `form:"kubernetes_resource,omitempty" json:"kubernetes_resource,omitempty" xml:"kubernetes_resource,omitempty"`
}

// UpdateRequestBody is the type of the "Backup Service" service "update"
// endpoint HTTP request body.
type UpdateRequestBody struct {
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	ID        *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Current state of the job
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Name of pachyderm instance backed up
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Namespace of resource backed up
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Name of target pod
	Pod *string `form:"pod,omitempty" json:"pod,omitempty" xml:"pod,omitempty"`
	// Name of container in pod
	Container *string `form:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// base64 encoded command to run in pod
	Command *string `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string `form:"storage_secret,omitempty" json:"storage_secret,omitempty" xml:"storage_secret,omitempty"`
	// base64 encoded json representation of object
	KubernetesResource *string `form:"kubernetes_resource,omitempty" json:"kubernetes_resource,omitempty" xml:"kubernetes_resource,omitempty"`
	BackupLocation     *string `form:"backup_location,omitempty" json:"backup_location,omitempty" xml:"backup_location,omitempty"`
}

// CreateResponseBody is the type of the "Backup Service" service "create"
// endpoint HTTP response body.
type CreateResponseBody struct {
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	ID        *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Current state of the job
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Name of pachyderm instance backed up
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Namespace of resource backed up
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Name of target pod
	Pod *string `form:"pod,omitempty" json:"pod,omitempty" xml:"pod,omitempty"`
	// Name of container in pod
	Container *string `form:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// base64 encoded command to run in pod
	Command *string `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string `form:"storage_secret,omitempty" json:"storage_secret,omitempty" xml:"storage_secret,omitempty"`
	// base64 encoded json representation of object
	KubernetesResource *string `form:"kubernetes_resource,omitempty" json:"kubernetes_resource,omitempty" xml:"kubernetes_resource,omitempty"`
	BackupLocation     *string `form:"backup_location,omitempty" json:"backup_location,omitempty" xml:"backup_location,omitempty"`
}

// GetResponseBody is the type of the "Backup Service" service "get" endpoint
// HTTP response body.
type GetResponseBody struct {
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	ID        *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Current state of the job
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Name of pachyderm instance backed up
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Namespace of resource backed up
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Name of target pod
	Pod *string `form:"pod,omitempty" json:"pod,omitempty" xml:"pod,omitempty"`
	// Name of container in pod
	Container *string `form:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// base64 encoded command to run in pod
	Command *string `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string `form:"storage_secret,omitempty" json:"storage_secret,omitempty" xml:"storage_secret,omitempty"`
	// base64 encoded json representation of object
	KubernetesResource *string `form:"kubernetes_resource,omitempty" json:"kubernetes_resource,omitempty" xml:"kubernetes_resource,omitempty"`
	BackupLocation     *string `form:"backup_location,omitempty" json:"backup_location,omitempty" xml:"backup_location,omitempty"`
}

// UpdateResponseBody is the type of the "Backup Service" service "update"
// endpoint HTTP response body.
type UpdateResponseBody struct {
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	ID        *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Current state of the job
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Name of pachyderm instance backed up
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Namespace of resource backed up
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Name of target pod
	Pod *string `form:"pod,omitempty" json:"pod,omitempty" xml:"pod,omitempty"`
	// Name of container in pod
	Container *string `form:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// base64 encoded command to run in pod
	Command *string `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string `form:"storage_secret,omitempty" json:"storage_secret,omitempty" xml:"storage_secret,omitempty"`
	// base64 encoded json representation of object
	KubernetesResource *string `form:"kubernetes_resource,omitempty" json:"kubernetes_resource,omitempty" xml:"kubernetes_resource,omitempty"`
	BackupLocation     *string `form:"backup_location,omitempty" json:"backup_location,omitempty" xml:"backup_location,omitempty"`
}

// DeleteResponseBody is the type of the "Backup Service" service "delete"
// endpoint HTTP response body.
type DeleteResponseBody struct {
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	ID        *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Current state of the job
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// Name of pachyderm instance backed up
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Namespace of resource backed up
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Name of target pod
	Pod *string `form:"pod,omitempty" json:"pod,omitempty" xml:"pod,omitempty"`
	// Name of container in pod
	Container *string `form:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// base64 encoded command to run in pod
	Command *string `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// Kubernetes secret containing S3 storage credentials
	StorageSecret *string `form:"storage_secret,omitempty" json:"storage_secret,omitempty" xml:"storage_secret,omitempty"`
	// base64 encoded json representation of object
	KubernetesResource *string `form:"kubernetes_resource,omitempty" json:"kubernetes_resource,omitempty" xml:"kubernetes_resource,omitempty"`
	BackupLocation     *string `form:"backup_location,omitempty" json:"backup_location,omitempty" xml:"backup_location,omitempty"`
}

// GetBackupNotFoundResponseBody is the type of the "Backup Service" service
// "get" endpoint HTTP response body for the "backup_not_found" error.
type GetBackupNotFoundResponseBody struct {
	// backup resource not found
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateBackupNotFoundResponseBody is the type of the "Backup Service" service
// "update" endpoint HTTP response body for the "backup_not_found" error.
type UpdateBackupNotFoundResponseBody struct {
	// backup resource not found
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteBackupNotFoundResponseBody is the type of the "Backup Service" service
// "delete" endpoint HTTP response body for the "backup_not_found" error.
type DeleteBackupNotFoundResponseBody struct {
	// backup resource not found
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "Backup Service" service.
func NewCreateRequestBody(p *backupservice.Backup) *CreateRequestBody {
	body := &CreateRequestBody{
		Name:               p.Name,
		Namespace:          p.Namespace,
		Pod:                p.Pod,
		Container:          p.Container,
		Command:            p.Command,
		StorageSecret:      p.StorageSecret,
		KubernetesResource: p.KubernetesResource,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "Backup Service" service.
func NewUpdateRequestBody(p *backupservice.Backupresult) *UpdateRequestBody {
	body := &UpdateRequestBody{
		CreatedAt:          p.CreatedAt,
		UpdatedAt:          p.UpdatedAt,
		DeletedAt:          p.DeletedAt,
		ID:                 p.ID,
		State:              p.State,
		Name:               p.Name,
		Namespace:          p.Namespace,
		Pod:                p.Pod,
		Container:          p.Container,
		Command:            p.Command,
		StorageSecret:      p.StorageSecret,
		KubernetesResource: p.KubernetesResource,
		BackupLocation:     p.BackupLocation,
	}
	return body
}

// NewCreateBackupresultAccepted builds a "Backup Service" service "create"
// endpoint result from a HTTP "Accepted" response.
func NewCreateBackupresultAccepted(body *CreateResponseBody) *backupserviceviews.BackupresultView {
	v := &backupserviceviews.BackupresultView{
		CreatedAt:          body.CreatedAt,
		UpdatedAt:          body.UpdatedAt,
		DeletedAt:          body.DeletedAt,
		ID:                 body.ID,
		State:              body.State,
		Name:               body.Name,
		Namespace:          body.Namespace,
		Pod:                body.Pod,
		Container:          body.Container,
		Command:            body.Command,
		StorageSecret:      body.StorageSecret,
		KubernetesResource: body.KubernetesResource,
		BackupLocation:     body.BackupLocation,
	}

	return v
}

// NewGetBackupresultOK builds a "Backup Service" service "get" endpoint result
// from a HTTP "OK" response.
func NewGetBackupresultOK(body *GetResponseBody) *backupserviceviews.BackupresultView {
	v := &backupserviceviews.BackupresultView{
		CreatedAt:          body.CreatedAt,
		UpdatedAt:          body.UpdatedAt,
		DeletedAt:          body.DeletedAt,
		ID:                 body.ID,
		State:              body.State,
		Name:               body.Name,
		Namespace:          body.Namespace,
		Pod:                body.Pod,
		Container:          body.Container,
		Command:            body.Command,
		StorageSecret:      body.StorageSecret,
		KubernetesResource: body.KubernetesResource,
		BackupLocation:     body.BackupLocation,
	}

	return v
}

// NewGetBackupNotFound builds a Backup Service service get endpoint
// backup_not_found error.
func NewGetBackupNotFound(body *GetBackupNotFoundResponseBody) *backupservice.BackupNotFound {
	v := &backupservice.BackupNotFound{
		Message: *body.Message,
	}

	return v
}

// NewUpdateBackupresultOK builds a "Backup Service" service "update" endpoint
// result from a HTTP "OK" response.
func NewUpdateBackupresultOK(body *UpdateResponseBody) *backupserviceviews.BackupresultView {
	v := &backupserviceviews.BackupresultView{
		CreatedAt:          body.CreatedAt,
		UpdatedAt:          body.UpdatedAt,
		DeletedAt:          body.DeletedAt,
		ID:                 body.ID,
		State:              body.State,
		Name:               body.Name,
		Namespace:          body.Namespace,
		Pod:                body.Pod,
		Container:          body.Container,
		Command:            body.Command,
		StorageSecret:      body.StorageSecret,
		KubernetesResource: body.KubernetesResource,
		BackupLocation:     body.BackupLocation,
	}

	return v
}

// NewUpdateBackupNotFound builds a Backup Service service update endpoint
// backup_not_found error.
func NewUpdateBackupNotFound(body *UpdateBackupNotFoundResponseBody) *backupservice.BackupNotFound {
	v := &backupservice.BackupNotFound{
		Message: *body.Message,
	}

	return v
}

// NewDeleteBackupresultOK builds a "Backup Service" service "delete" endpoint
// result from a HTTP "OK" response.
func NewDeleteBackupresultOK(body *DeleteResponseBody) *backupserviceviews.BackupresultView {
	v := &backupserviceviews.BackupresultView{
		CreatedAt:          body.CreatedAt,
		UpdatedAt:          body.UpdatedAt,
		DeletedAt:          body.DeletedAt,
		ID:                 body.ID,
		State:              body.State,
		Name:               body.Name,
		Namespace:          body.Namespace,
		Pod:                body.Pod,
		Container:          body.Container,
		Command:            body.Command,
		StorageSecret:      body.StorageSecret,
		KubernetesResource: body.KubernetesResource,
		BackupLocation:     body.BackupLocation,
	}

	return v
}

// NewDeleteBackupNotFound builds a Backup Service service delete endpoint
// backup_not_found error.
func NewDeleteBackupNotFound(body *DeleteBackupNotFoundResponseBody) *backupservice.BackupNotFound {
	v := &backupservice.BackupNotFound{
		Message: *body.Message,
	}

	return v
}

// ValidateGetBackupNotFoundResponseBody runs the validations defined on
// get_backup_not_found_response_body
func ValidateGetBackupNotFoundResponseBody(body *GetBackupNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateBackupNotFoundResponseBody runs the validations defined on
// update_backup_not_found_response_body
func ValidateUpdateBackupNotFoundResponseBody(body *UpdateBackupNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDeleteBackupNotFoundResponseBody runs the validations defined on
// delete_backup_not_found_response_body
func ValidateDeleteBackupNotFoundResponseBody(body *DeleteBackupNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}
