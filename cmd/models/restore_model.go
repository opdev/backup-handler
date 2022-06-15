package models

import (
	"context"
	"database/sql"
	"errors"
	"log"

	restoreservice "github.com/opdev/backup-handler/gen/restore_service"
)

// CreateRestore stores a new restore request
func CreateRestore(restore *restoreservice.Restoreresult) error {
	db, err := connect()
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return createRestore(context.Background(), db, restore)
}

func createRestore(ctx context.Context, db *sql.DB, restore *restoreservice.Restoreresult) error {
	query := "INSERT INTO restores(created_at, id, name, namespace, backup_location, storage_secret) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer db.Close()

	results, err := stmt.ExecContext(ctx,
		restore.CreatedAt,
		restore.ID,
		restore.Name,
		restore.Namespace,
		restore.BackupLocation,
		restore.StorageSecret,
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

// GetRestore returns a restore from the database
func GetRestore(restore *restoreservice.Restoreresult) error {
	db, err := connect()
	if err != nil {
		return err
	}

	return getRestore(context.Background(), db, restore)
}

func getRestore(ctx context.Context, db *sql.DB, restore *restoreservice.Restoreresult) error {
	query := "SELECT created_at, updated_at, deleted_at, id, name, namespace, backup_location, storage_secret, kubernetes_object, db FROM restores WHERE id = ?"
	row := db.QueryRowContext(ctx, query, restore.ID)

	if err := row.Scan(
		&restore.CreatedAt,
		&restore.UpdatedAt,
		&restore.DeletedAt,
		&restore.ID,
		&restore.Name,
		&restore.Namespace,
		&restore.BackupLocation,
		&restore.StorageSecret,
		&restore.KubernetesResource,
		&restore.Database,
	); err != nil {
		log.Printf("error reading responses; %+v\n", err)
		return err
	}

	return nil
}

// UpdateRestore updates a restore object in the database
func UpdateRestore(restore *restoreservice.Restoreresult) (int64, error) {
	db, err := connect()
	if err != nil {
		return 0, err
	}

	return updateRestore(context.Background(), db, restore)
}

func updateRestore(ctx context.Context, db *sql.DB, restore *restoreservice.Restoreresult) (int64, error) {
	query := `UPDATE restores SET updated_at = ?, backup_location = ?, storage_secret = ?, kubernetes_object = ?, db = ? WHERE id = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	defer stmt.Close()

	results, err := stmt.ExecContext(ctx,
		restore.UpdatedAt,
		restore.BackupLocation,
		restore.StorageSecret,
		restore.KubernetesResource,
		restore.Database,
		restore.ID,
	)
	if err != nil {
		return 0, err
	}

	return results.RowsAffected()
}

// DeleteRestore updates a restore object in the database
func DeleteRestore(restore *restoreservice.Restoreresult) (int64, error) {
	db, err := connect()
	if err != nil {
		return 0, err
	}

	return deleteRestore(context.Background(), db, restore)
}

func deleteRestore(ctx context.Context, db *sql.DB, restore *restoreservice.Restoreresult) (int64, error) {
	query := `UPDATE restores SET updated_at = ?, deleted_at = ?, backup_location = ?, storage_secret = ?, kubernetes_object = ?, db = ? WHERE id = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	defer stmt.Close()

	results, err := stmt.ExecContext(ctx,
		restore.UpdatedAt,
		restore.DeletedAt,
		restore.BackupLocation,
		restore.StorageSecret,
		restore.KubernetesResource,
		restore.Database,
		restore.ID,
	)
	if err != nil {
		return 0, err
	}

	return results.RowsAffected()
}
