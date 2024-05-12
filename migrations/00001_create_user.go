package migrations

import (
	"app/database"
	"app/models"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUser, downCreateUser)
}

func upCreateUser(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.CreateTable(&models.User{})
}

func downCreateUser(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.DropTable(&models.User{})
}
