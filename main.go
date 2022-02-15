package main

import (
	"log"

	"github.com/opdev/backup-handler/pkg/command"
	"github.com/opdev/backup-handler/pkg/restapi"
)

func main() {
	// Start the REST api in a goroutine
	go func() {
		if err := restapi.Start(); err != nil {
			log.Panic("error starting REST api", err)
		}

	}()

	command.BackupDispatch()
}
