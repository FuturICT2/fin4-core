package dbservice

import (
	"database/sql"
	"fmt"
	"log"

	// we need to import mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/FuturICT2/fin4-core/server/env"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

// DB database connection type
type DB struct {
	*sql.DB
}

//MustConnect connects to database or panics
func MustConnect(dataSourceName string) *DB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("db:MustConnect:cant open sql")
		log.Fatal(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("db:MustConnect:db unreachable")
		log.Fatal(err.Error())
	}
	if err := migrateDB(db); err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("db:MustConnect:unable to migrate")
		log.Fatal("migrateDB: ", err)
	}
	return &DB{db}
}

// TestConnect connects to test database
func TestConnect() *DB {
	dbName := "exchange_test"
	db := MustConnect(
		env.MustGetenv("TEST_DATA_SOURCE_NAME"),
	)
	db.Exec(fmt.Sprintf("DROP DATABASE %s IF EXISTS", dbName))
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE  IF NOT EXISTS `%s`", dbName))
	if err != nil {
		log.Fatal("Can't create test database")
	}
	if err := migrateDB(db.DB); err != nil {
		log.Fatal("migrateDB: ", err)
	}
	return db
}

// helper function to setup the databsae by performing automated database migration steps.
func migrateDB(db *sql.DB) error {
	var migrations = &migrate.FileMigrationSource{
		Dir: env.MustGetenv("DB_MIGRATIONS_PATH"),
	}
	_, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	return err
}
