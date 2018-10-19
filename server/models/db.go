package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // driver

	"github.com/FuturICT2/fin4-core/server/util"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

// ID Database id type
type ID uint64

// DbBool TINYINT
type DbBool int8

const (
	// DbTrue a tinyint true
	DbTrue = 1
	// DbFalse a tinyint False
	DbFalse = 0
)

// DB database connection type
type DB struct {
	*sql.DB
}

//MustConnect connects to database or panics
func MustConnect(dataSourceName string) *DB {
	log.Println("Connecting to the database: ", dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("db:MustConnect:e1")
		log.Fatal(err.Error())
	}
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("db:MustConnect:e2")
		log.Fatal(err.Error())
	}
	if err := migrateDB(db); err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("db:MustConnect:e3")
		log.Fatal("migrateDB: ", err)
	}
	return &DB{db}
}

// TestConnect connects to test database
func TestConnect() *DB {
	dbName := "fin4_test"
	dsn := fmt.Sprintf(
		`%s:%s@/%s?charset=utf8mb4,utf8&parseTime=true"`,
		util.MustGetenv("TEST_DB_USER"),
		util.MustGetenv("TEST_DB_PASS"),
		dbName,
	)
	db := MustConnect(dsn)
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
		Dir: util.MustGetenv("DB_MIGRATIONS_PATH"),
	}
	_, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	return err
}

// DBInterface db interface data type
type DBInterface interface {
	NewUserModel() UserStore
}

// UserModel struct holding user model
type UserModel struct {
	*DB
}

// NewUserModel creates a new user model
func (db *DB) NewUserModel() UserStore {
	return &UserModel{db}
}
