package command

import (
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	backupv1 "github.com/opdev/backup-handler/api/v1"
	"github.com/walle/targz"
	"k8s.io/apimachinery/pkg/types"
)

func StartRestore(restore *backupv1.Restore) error {
	// backupKey refers to the key of the object to be
	// restored from the S3 storage bucket
	asset := path.Join(
		"/",
		"tmp",
		path.Base(restore.Backup),
	)

	size, err := downloadBackup(restore, asset)
	if err != nil {
		return err
	}

	log.Printf("downloaded %d bytes of backup to %s.\n", size, asset)
	if err := targz.Extract(asset, "/tmp"); err != nil {
		return err
	}

	return nil
}

func downloadBackup(restore *backupv1.Restore, backupFile string) (int64, error) {
	creds, err := getUploadCredentials(
		types.NamespacedName{
			Name:      restore.StorageSecret,
			Namespace: restore.Metadata.Namespace,
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
