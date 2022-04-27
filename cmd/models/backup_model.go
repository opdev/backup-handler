package models

import (
	"context"
	"database/sql"
	"errors"

	backupservice "github.com/opdev/backup-handler/gen/backup_service"
)

func CreateBackup(backup *backupservice.Backupresult) error {
	db, err := connect()
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return createBackup(context.Background(), db, backup)
}

func createBackup(ctx context.Context, db *sql.DB, backup *backupservice.Backupresult) error {
	query := `INSERT INTO backups(created_at, id, name, namespace, state, pod, container,
command, storage_secret, kube_resource) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer db.Close()

	results, err := stmt.ExecContext(ctx,
		backup.CreatedAt,
		backup.ID,
		backup.Name,
		backup.Namespace,
		backup.State,
		backup.Pod,
		backup.Container,
		backup.Command,
		backup.StorageSecret,
		backup.KubernetesResource,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	count, err := results.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("error adding record")
	}

	return nil
}

// GetBackup returns a specific backup objecy=t
func GetBackup(backup *backupservice.Backupresult) error {
	db, err := connect()
	if err != nil {
		return err
	}

	return getBackup(context.Background(), db, backup)
}

func getBackup(ctx context.Context, db *sql.DB, backup *backupservice.Backupresult) error {
	query := `SELECT created_at, updated_at, deleted_at, id, name, namespace, state, pod, container,
command, backup_location, storage_secret, kube_resource FROM backups WHERE id = ?`
	var location sql.NullString
	row := db.QueryRowContext(ctx, query, backup.ID)
	if err := row.Scan(
		&backup.CreatedAt,
		&backup.UpdatedAt,
		&backup.DeletedAt,
		&backup.ID,
		&backup.Name,
		&backup.Namespace,
		&backup.State,
		&backup.Pod,
		&backup.Container,
		&backup.Command,
		&location,
		&backup.StorageSecret,
		&backup.KubernetesResource,
	); err != nil {
		return err
	}

	// Bug Fix: use location place holder since backup.UploadLocation could be NULL
	// This is needed since the database/sql package can not handle NULL values
	if location.Valid {
		backup.BackupLocation = &location.String
	}

	return nil
}

// UpdateBackup updates the properties of a backup resource
func UpdateBackup(backup *backupservice.Backupresult) (int64, error) {
	db, err := connect()
	if err != nil {
		return 0, err
	}

	if err := db.Ping(); err != nil {
		return 0, err
	}

	return updateBackup(context.Background(), db, backup)
}

// TODO: cleanup function implementation
func updateBackup(ctx context.Context, db *sql.DB, backup *backupservice.Backupresult) (int64, error) {
	query := `UPDATE backups SET updated_at = ?, name = ?,  state = ?,  pod = ?,  container = ?,
command = ?, backup_location = ?, storage_secret = ?, kube_resource = ? WHERE id = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	defer stmt.Close()

	results, err := stmt.ExecContext(ctx,
		backup.UpdatedAt,
		backup.Name,
		backup.State,
		backup.Pod,
		backup.Container,
		backup.Command,
		backup.BackupLocation,
		backup.StorageSecret,
		backup.KubernetesResource,
		backup.ID,
	)
	if err != nil {
		return 0, err
	}

	return results.RowsAffected()
}

func DeleteBackup(backup *backupservice.Backupresult) (int64, error) {
	db, err := connect()
	if err != nil {
		return 0, err
	}

	if err := db.Ping(); err != nil {
		return 0, err
	}

	return deleteBackup(context.Background(), db, backup)
}

func deleteBackup(ctx context.Context, db *sql.DB, backup *backupservice.Backupresult) (int64, error) {
	query := `UPDATE backups SET deleted_at = ? WHERE id = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	defer stmt.Close()

	results, err := stmt.ExecContext(ctx,
		backup.DeletedAt,
		backup.ID,
	)
	if err != nil {
		return 0, err
	}

	return results.RowsAffected()
}
