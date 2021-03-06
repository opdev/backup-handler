// Code generated by goa v3.7.3, DO NOT EDIT.
//
// Restore Service views
//
// Command:
// $ goa gen github.com/opdev/backup-handler/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Restoreresult is the viewed result type that is projected based on a view.
type Restoreresult struct {
	// Type to project
	Projected *RestoreresultView
	// View to render
	View string
}

// RestoreresultView is a type that runs validations on a projected type.
type RestoreresultView struct {
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

var (
	// RestoreresultMap is a map indexing the attribute names of Restoreresult by
	// view name.
	RestoreresultMap = map[string][]string{
		"default": {
			"created_at",
			"updated_at",
			"deleted_at",
			"id",
			"name",
			"namespace",
			"backup_location",
			"storage_secret",
			"kubernetes_resource",
			"database",
		},
	}
)

// ValidateRestoreresult runs the validations defined on the viewed result type
// Restoreresult.
func ValidateRestoreresult(result *Restoreresult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateRestoreresultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateRestoreresultView runs the validations defined on RestoreresultView
// using the "default" view.
func ValidateRestoreresultView(result *RestoreresultView) (err error) {

	return
}
