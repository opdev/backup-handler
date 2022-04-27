package backuphandler

import (
	"time"

	"github.com/google/uuid"
)

func utcTime() *string {
	t := time.Now().UTC().String()
	return &t
}

func genUUID() *string {
	id := uuid.New().String()
	return &id
}
