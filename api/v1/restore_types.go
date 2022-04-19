package v1

import "github.com/google/uuid"

// Restore object
type Restore struct {
	Metadata `json:"metadata,omitempty"`
	// Name of the backup object in S3 bucket to be restored
	Backup string `json:"backup,omitempty"`
	// Destination refers to name and namespace to which the
	// backup object would be restored to
	Destination `json:"destination,omitempty"`
}

type Destination struct {
	// Name of the destination of the backup restore
	Name string `json:"name,omitempty"`
	// Namespace ion which to restore pachyderm backup
	Namespace string `json:"namespace,omitempty"`
}

func (r *Restore) New() {
	r.Metadata.CreatedAt = utcTime()
	r.Metadata.ID = uuid.New()
}

func (r *Restore) SetUpdatedTime() {
	r.Metadata.UpdatedAt = utcTime()
}

func (r *Restore) SetDeletedTime() {
	r.Metadata.DeletedAt = utcTime()
}
