package command

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	backupv1 "github.com/opdev/backup-handler/api/v1"
	"k8s.io/apimachinery/pkg/types"
)

func StartRestore(restore *backupv1.Restore) error {

	if err := downloadBackup(restore); err != nil {
		return err
	}

	return nil
}

func downloadBackup(restore *backupv1.Restore) error {
	creds, err := getUploadCredentials(
		types.NamespacedName{
			Name:      restore.StorageSecret,
			Namespace: restore.Metadata.Namespace,
		},
	)
	if err != nil {
		return err
	}

	_session, err := creds.awsSession()
	if err != nil {
		return err
	}

	downloader := s3manager.NewDownloader(_session)

	// Create file to write S3 object contents to
	backup, err := os.Create("/tmp/restore.tar.gz")
	if err != nil {
		return err
	}

	// write contents of the S3 object to the file
	n, err := downloader.Download(backup, &s3.GetObjectInput{
		Bucket: aws.String(creds.bucket),
		Key:    aws.String(restore.Backup),
	})
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes %s file downloaded!", n, backup.Name())

	return nil
}
