package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/dentych/dinner-dash/internal/config"
)

var db *sqlx.DB

var ConnectionString string

// Init will setup a new database connection. The method will panic
// if a database connection can not be established.
func Init(config config.DatabaseConfig) {
	format := "host=%s user=%s password=%s dbname=%s sslmode=disable"
	ConnectionString = fmt.Sprintf(format, config.Hostname, config.Username, config.Password, config.Database)
	db = sqlx.MustConnect("postgres", ConnectionString)
}

// GetConnection will return a database object, which can be used to perform queries.
func GetConnection() *sqlx.DB {
	if db == nil {
		//logging.Error.Fatal("Database is nil. You must initialize before getting the connection.")
	}

	return db
}

func GetConnectionString() string {
	if ConnectionString == "" {
		//logging.Error.Fatal("You must initialize the database before calling GetConnectionString")
	}

	return ConnectionString
}