package models

import (
	"context"
	"database/sql"
	"errors"

	backupv1 "github.com/opdev/backup-handler/api/v1"
)

func CreateRestore(restore *backupv1.Restore) error {
	db, err := connect()
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	restore.New()
	return createRestore(context.Background(), db, restore)
}

func createRestore(ctx context.Context, db *sql.DB, restore *backupv1.Restore) error {
	query := "INSERT INTO restores(created_at, id, name, namespace, backup_name) VALUES(?, ?, ?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer db.Close()

	results, err := stmt.ExecContext(ctx,
		restore.CreatedAt,
		restore.ID,
		restore.Destination.Name,
		restore.Destination.Namespace,
		restore.Backup,
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

func GetRestore(restore *backupv1.Restore) error {
	db, err := connect()
	if err != nil {
		return err
	}

	return getRestore(context.Background(), db, restore)
}

func getRestore(ctx context.Context, db *sql.DB, restore *backupv1.Restore) error {
	query := "SELECT created_at, id, name, namespace, backup_name FROM restores WHERE id = ?"
	row := db.QueryRowContext(ctx, query, restore.ID)

	if err := row.Scan(
		&restore.CreatedAt,
		&restore.UpdatedAt,
		&restore.DeletedAt,
		&restore.ID,
		&restore.Destination.Name,
		&restore.Destination.Namespace,
		&restore.Backup,
	); err != nil {
		return err
	}

	return nil
}
