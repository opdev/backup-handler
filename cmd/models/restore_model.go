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
	query := "INSERT INTO restores(created_at, id, name, namespace, backup_name, storage_secret) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer db.Close()

	results, err := stmt.ExecContext(ctx,
		restore.CreatedAt,
		restore.ID,
		restore.DestinationName,
		restore.DestinationNamespace,
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
	query := "SELECT created_at, updated_at, deleted_at, id, name, namespace, backup_name, storage_secret FROM restores WHERE id = ?"
	row := db.QueryRowContext(ctx, query, restore.ID)

	if err := row.Scan(
		&restore.CreatedAt,
		&restore.UpdatedAt,
		&restore.DeletedAt,
		&restore.ID,
		&restore.DestinationName,
		&restore.DestinationNamespace,
		&restore.BackupLocation,
		&restore.StorageSecret,
	); err != nil {
		log.Println("error reading responses")
		return err
	}

	return nil
}
