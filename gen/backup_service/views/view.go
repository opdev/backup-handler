// Code generated by goa v3.7.3, DO NOT EDIT.
//
// Backup Service views
//
// Command:
// $ goa gen github.com/opdev/backup-handler/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Backupresult is the viewed result type that is projected based on a view.
type Backupresult struct {
	// Type to project
	Projected *BackupresultView
	// View to render
	View string
}

// BackupresultView is a type that runs validations on a projected type.
type BackupresultView struct {
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

var (
	// BackupresultMap is a map indexing the attribute names of Backupresult by
	// view name.
	BackupresultMap = map[string][]string{
		"default": {
			"created_at",
			"updated_at",
			"deleted_at",
			"id",
			"state",
			"name",
			"namespace",
			"pod",
			"container",
			"command",
			"storage_secret",
			"kubernetes_resource",
			"location",
		},
	}
)

// ValidateBackupresult runs the validations defined on the viewed result type
// Backupresult.
func ValidateBackupresult(result *Backupresult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateBackupresultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateBackupresultView runs the validations defined on BackupresultView
// using the "default" view.
func ValidateBackupresultView(result *BackupresultView) (err error) {
	if result.State != nil {
		if !(*result.State == "queued" || *result.State == "running" || *result.State == "completed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.state", *result.State, []interface{}{"queued", "running", "completed"}))
		}
	}
	return
}
