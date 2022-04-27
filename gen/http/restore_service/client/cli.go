// Code generated by goa v3.7.2, DO NOT EDIT.
//
// Restore Service HTTP client CLI support package
//
// Command:
// $ goa gen github.com/opdev/backup-handler/design

package client

import (
	"encoding/json"
	"fmt"

	restoreservice "github.com/opdev/backup-handler/gen/restore_service"
)

// BuildCreatePayload builds the payload for the Restore Service create
// endpoint from CLI flags.
func BuildCreatePayload(restoreServiceCreateBody string) (*restoreservice.Restore, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(restoreServiceCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"backup_location\": \"pachyderm-backup.tar.gz\",\n      \"destination_name\": \"pachyderm-restore\",\n      \"destination_namespace\": \"ai-namespace\",\n      \"name\": \"pachdyderm-sample\",\n      \"namespace\": \"testing\",\n      \"storage_secret\": \"example-aws-secret\"\n   }'")
		}
	}
	v := &restoreservice.Restore{
		Name:                 body.Name,
		Namespace:            body.Namespace,
		StorageSecret:        body.StorageSecret,
		DestinationName:      body.DestinationName,
		DestinationNamespace: body.DestinationNamespace,
		BackupLocation:       body.BackupLocation,
	}

	return v, nil
}

// BuildGetPayload builds the payload for the Restore Service get endpoint from
// CLI flags.
func BuildGetPayload(restoreServiceGetID string) (*restoreservice.GetPayload, error) {
	var id string
	{
		id = restoreServiceGetID
	}
	v := &restoreservice.GetPayload{}
	v.ID = &id

	return v, nil
}

// BuildUpdatePayload builds the payload for the Restore Service update
// endpoint from CLI flags.
func BuildUpdatePayload(restoreServiceUpdateBody string) (*restoreservice.Restoreresult, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(restoreServiceUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"backup_location\": \"pachyderm-backup.tar.gz\",\n      \"created_at\": \"2019-10-12 07:20:50.52\",\n      \"database\": \"database.sql\",\n      \"deleted_at\": \"2019-10-12 07:20:54.52\",\n      \"destination_name\": \"pachyderm-restore\",\n      \"destination_namespace\": \"ai-namespace\",\n      \"id\": \"00000-090000-0000000-000000\",\n      \"kubernetes_resource\": \"pachyderm.yaml\",\n      \"name\": \"pachdyderm-sample\",\n      \"namespace\": \"testing\",\n      \"storage_secret\": \"example-aws-secret\",\n      \"updated_at\": \"2019-10-12 07:20:52.52\"\n   }'")
		}
	}
	v := &restoreservice.Restoreresult{
		CreatedAt:            body.CreatedAt,
		UpdatedAt:            body.UpdatedAt,
		DeletedAt:            body.DeletedAt,
		ID:                   body.ID,
		Name:                 body.Name,
		Namespace:            body.Namespace,
		BackupLocation:       body.BackupLocation,
		DestinationName:      body.DestinationName,
		DestinationNamespace: body.DestinationNamespace,
		StorageSecret:        body.StorageSecret,
		KubernetesResource:   body.KubernetesResource,
		Database:             body.Database,
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the Restore Service delete
// endpoint from CLI flags.
func BuildDeletePayload(restoreServiceDeleteID string) (*restoreservice.DeletePayload, error) {
	var id string
	{
		id = restoreServiceDeleteID
	}
	v := &restoreservice.DeletePayload{}
	v.ID = &id

	return v, nil
}
