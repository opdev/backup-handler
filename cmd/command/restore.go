package command

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/opdev/backup-handler/cmd/models"
	restoreservice "github.com/opdev/backup-handler/gen/restore_service"
	"github.com/walle/targz"
	"k8s.io/apimachinery/pkg/types"
)

// StartRestore command is the invoked to start restoring from backup
func StartRestore(restore *restoreservice.Restoreresult, logger *log.Logger) error {
	// backupKey refers to the key of the object to be
	// restored from the S3 storage bucket
	asset := path.Join(
		"/",
		"tmp",
		path.Base(*restore.BackupLocation),
	)

	size, err := downloadBackup(restore, asset)
	if err != nil {
		return err
	}

	logger.Printf("%d bytes of %s downloaded.\n", size, path.Base(asset))
	if err := targz.Extract(asset, "/tmp"); err != nil {
		return err
	}

	if err := loadBackup(asset, restore); err != nil {
		return err
	}

	// TODO: store the payload to a database awaiting
	// api request for status
	now := time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
	restore.UpdatedAt = &now
	changes, err := models.UpdateRestore(restore)
	if err != nil {
		return err
	}
	if changes == 1 {
		logger.Println("backup extracted to /tmp")
	}

	return nil
}

func downloadBackup(restore *restoreservice.Restoreresult, backupFile string) (int64, error) {
	creds, err := getUploadCredentials(
		types.NamespacedName{
			Name:      *restore.StorageSecret,
			Namespace: *restore.Namespace,
		},
	)
	if err != nil {
		return 0, err
	}

	_session, err := creds.awsSession()
	if err != nil {
		return 0, err
	}

	downloader := s3manager.NewDownloader(_session)

	// Create file to write S3 object contents to
	backup, err := os.Create(backupFile)
	if err != nil {
		return 0, err
	}
	defer backup.Close()

	// write contents of the S3 object to the file
	return downloader.Download(backup, &s3.GetObjectInput{
		Bucket: aws.String(creds.bucket),
		Key: aws.String(
			path.Join(
				"backups",
				path.Base(backupFile),
			),
		),
	})
}

// Load backup takes the backup directory and restore as parameters
// updates the restore object with the database dump and custom resource
// and returns an error
func loadBackup(backup string, restore *restoreservice.Restoreresult) error {
	// Split the string at the periods
	// and get the first item in the slice
	dir := strings.Split(backup, ".")[0]

	cr, err := readFileContents(path.Join(dir, "cr.json"))
	if err != nil {
		return err
	}

	db, err := readFileContents(path.Join(dir, "database.tar"))
	if err != nil {
		return err
	}

	resource := base64.StdEncoding.EncodeToString(cr)
	database := base64.StdEncoding.EncodeToString(db)
	restore.KubernetesResource = &resource
	restore.Database = &database

	return nil
}

func readFileContents(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
