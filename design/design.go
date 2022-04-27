package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("Backup Handler", func() {
	Title("Backup Handler Service")
	Description("A service to handle backup for Pachyderm object")

	Error("backup_not_found", func() {
		Description("This error is returned when backup does not exist")
		Fault()
	})
	HTTP(func() {
		Response("backup_not_found", StatusNotFound)
	})

	Server("http", func() {
		Host("localhost", func() {
			URI("http://localhost:8890")
		})

	})
})

// Backup resource is used to make a request to
// create a new backup object
var Backup = Type("Backup", func() {
	Attribute("name", String, func() {
		Description("Name of pachyderm instance backed up")
		Example("pachdyderm-sample")
	})
	Attribute("namespace", String, func() {
		Description("Namespace of resource backed up")
		Example("testing")
	})
	Attribute("pod", String, func() {
		Description("Name of target pod")
		Example("pachd-65734594-sdg46")
	})
	Attribute("container", String, func() {
		Description("Name of container in pod")
		Example("pachd")
	})
	Attribute("command", String, func() {
		Description("base64 encoded command to run in pod")
		Example("hostname -f")
	})
	Attribute("storage_secret", String, func() {
		Description("Kubernetes secret containing S3 storage credentials")
		Example("example-aws-secret")
	})
	Attribute("kubernetes_resource", String, func() {
		Description("base64 encoded json representation of object")
		Example(`{"kind":"Pachyderm",...}`)
	})
})

// BackupResult is the Backup representation returned by the API server
var BackupResult = ResultType("BackupResult", func() {
	Reference(Backup)
	Attributes(func() {
		Attribute("created_at", String, func() {
			Example("2019-10-12 07:20:50.52")
		})
		Attribute("updated_at", String, func() {
			Example("2019-10-12 07:20:52.52")
		})
		Attribute("deleted_at", String, func() {
			Example("2019-10-12 07:20:54.52")
		})
		Attribute("id", String, func() {
			Example("00000-090000-0000000-000000")
		})
		Attribute("state", String, func() {
			Description("Current state of the job")
			Enum("queued", "running", "completed")
		})
		Attribute("name")
		Attribute("namespace")
		Attribute("pod")
		Attribute("container")
		Attribute("command")
		Attribute("storage_secret")
		Attribute("kubernetes_resource")
		Attribute("backup_location", String, func() {
			Example("http://minio.local/backups/pachyderm-backup.tar.gz")
		})
	})
})

// Restore is a representation of the request to restore a pachyderm instance
var Restore = Type("Restore", func() {
	Attribute("name", String, func() {
		Description("Name of pachyderm instance to restore to")
		Example("pachdyderm-sample")
	})
	Attribute("namespace", String, func() {
		Description("Namespace to restore to")
		Example("testing")
	})
	Attribute("storage_secret", String, func() {
		Description("Kubernetes secret containing S3 storage credentials")
		Example("example-aws-secret")
	})
	Attribute("destination_name", String, func() {
		Description("name of pachyderm instance to restore to")
		Example("pachyderm-restore")
	})
	Attribute("destination_namespace", String, func() {
		Description("namespace to restore pachyderm to")
		Example("ai-namespace")
	})
	Attribute("backup_location", String, func() {
		Example("pachyderm-backup.tar.gz")
	})
})

// RestoreResult is a representation of the restore request returned by the API server
var RestoreResult = ResultType("RestoreResult", func() {
	Reference(Restore)
	Attributes(func() {
		Attribute("created_at", String, func() {
			Example("2019-10-12 07:20:50.52")
		})
		Attribute("updated_at", String, func() {
			Example("2019-10-12 07:20:52.52")
		})
		Attribute("deleted_at", String, func() {
			Example("2019-10-12 07:20:54.52")
		})
		Attribute("id", String, func() {
			Example("00000-090000-0000000-000000")
		})
		Attribute("name")
		Attribute("namespace")
		Attribute("backup_location")
		Attribute("destination_name")
		Attribute("destination_namespace")
		Attribute("storage_secret")
		Attribute("kubernetes_resource", String, func() {
			Description("base64 encoded kubernetes object")
			Example("pachyderm.yaml")
		})
		Attribute("database", String, func() {
			Description("base64 encoded database dump")
			Example("database.sql")
		})
	})
})

var _ = Service("Backup Service", func() {
	Description("Service to handle backup requests")

	Method("create", func() {
		Description("New backup request")
		Payload(Backup)
		Result(BackupResult, StatusAccepted)
		HTTP(func() {
			POST("/backups")
			Response(StatusAccepted)
		})
	})

	Method("get", func() {
		Description("Obtain backup request")
		Payload(func() {
			Reference(BackupResult)
			Attribute("id")
		})
		Result(BackupResult)
		HTTP(func() {
			GET("/backups/{id}")
			Response(StatusOK)
			Response("backup_not_found", StatusNotFound, func() {
				Description("Error is returned when backup is not found")
			})
		})
		Error("backup_not_found", BackupNotFound, "backup not found")
	})

	Method("update", func() {
		Description("Update backup request")
		Payload(BackupResult)
		Result(BackupResult)
		HTTP(func() {
			PUT("/backups")
			Response(StatusOK)
			Response("backup_not_found", StatusNotFound, func() {
				Description("Error is returned when backup is not found")
			})
		})
		Error("backup_not_found", BackupNotFound, "backup not found")
	})

	Method("delete", func() {
		Description("Mark complete backup request")
		Payload(func() {
			Reference(BackupResult)
			Attribute("id")
		})
		Result(BackupResult)
		HTTP(func() {
			DELETE("/backups/{id}")
			Response(StatusOK)
			Response("backup_not_found", StatusNotFound, func() {
				Description("Error is returned when backup is not found")
			})
		})
		Error("backup_not_found", BackupNotFound, "backup not found")
	})
})

var _ = Service("Restore Service", func() {
	Description("Service to handle restore requests")

	Method("create", func() {
		Description("New restore request")
		Payload(Restore)
		Result(RestoreResult)
		HTTP(func() {
			POST("/restores")
			Response(StatusOK)
			Response("backup_not_found", StatusNotFound, func() {
				Description("Error is returned when backup location is not found")
			})
		})
		Error("backup_not_found", BackupNotFound, "backup not found")
	})

	Method("get", func() {
		Description("Get restore request")
		Payload(func() {
			Reference(RestoreResult)
			Attribute("id")
		})
		Result(RestoreResult)
		HTTP(func() {
			GET("/restores/{id}")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update restore request")
		Payload(RestoreResult)
		Result(RestoreResult)
		HTTP(func() {
			PUT("/restores")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Mark complete restore request")
		Payload(func() {
			Reference(RestoreResult)
			Attribute("id")
		})
		Result(RestoreResult)
		HTTP(func() {
			DELETE("/restores/{id}")
			Response(StatusOK)
		})
	})

	Files("/openapi3.json", "./gen/http/openapi3.json")
})

// BackupNotFound is returned when either the backup ID or location is invalid
var BackupNotFound = Type("BackupNotFound", func() {
	Description("Backup not found error is returned when the backup is not found.")
	Field(1, "message", String, "backup resource not found")
	Required("message")
})
