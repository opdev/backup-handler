package backuphandler

import (
	"time"

	"github.com/google/uuid"
)

func utcTime() *string {
	t := time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
	return &t
}

func genUUID() *string {
	id := uuid.New().String()
	return &id
}
