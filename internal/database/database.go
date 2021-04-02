package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/dentych/dinner-dash/internal/config"
)

// Init will setup a new database connection. The method will panic
// if a database connection can not be established.
func Init(config config.DatabaseConfig) *sqlx.DB {
	format := "host=%s user=%s password=%s dbname=%s sslmode=disable"
	connectionString := fmt.Sprintf(format, config.Hostname, config.Username, config.Password, config.Database)
	db := sqlx.MustOpen("postgres", connectionString)
	return db
}
