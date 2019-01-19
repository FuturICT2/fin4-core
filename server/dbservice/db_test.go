package dbservice_test

import (
	"fmt"
	"log"
	"time"

	"github.com/FuturICT2/fin4-core/exchange/server/appstrings"
	"github.com/FuturICT2/fin4-core/server/dbservice"
	"github.com/FuturICT2/fin4-core/server/env"
)

// Default test data

var testDbTables = []string{
	"crypto_deposit",
	"crypto_user_deposit_address",
	"user_password_reset",
	"user_deposit",
	"user_change_email_confirm",
	"user_balance",
	"trade_transaction",
	"trade_order",
	"trade_market",
	"trade_asset",
	"user",
}

var testDbAssets = []struct {
	ID     int
	Name   string
	Symbol string
}{
	{1, "Bitcoin", "BTC"},
	{2, "Ethereum", "ETH"},
	{3, "USD", "USD"},
	{4, "Euro", "EUR"},
}

var testDbMarkets = []struct {
	ID          int
	AssetID     int
	BaseAssetID int
	IsActive    bool
}{
	{1, 2, 1, true},
	{2, 3, 1, true},
	{3, 4, 1, true},
}

var testDbUsers = []struct {
	ID    int
	Email string
}{
	{1, "e1@t.c"},
	{2, "e2@t.c"},
	{3, "e3@t.c"},
	{4, "e4@t.c"},
	{5, "e5@t.c"},
}

func SetupTestDB() *dbservice.DB {
	db := dbservice.MustConnect(env.MustGetenv("TEST_DATA_SOURCE_NAME"))
	ResetDB(db)
	return db
}

func TearDownTestDB(db *dbservice.DB) {
	for _, tableName := range testDbTables {
		truncateTable(db, tableName)
	}
	db.Close()
}

func ResetDB(db *dbservice.DB) {
	for _, tableName := range testDbTables {
		truncateTable(db, tableName)
	}
	createTestUsers(db)
	createTestAssets(db)
	createTestMarkets(db)
}

func createTestUsers(db *dbservice.DB) {
	stmt, err := db.Prepare(`INSERT INTO user SET
		id = ?,
		email = ?,
		password = ?,
		salt = ?,
		createdAt = ?,
		updatedAt = ?,
		agreeToTerms = ?`,
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("createTestUsers:1 %s", err.Error()))
	}
	for _, entry := range testDbUsers {
		_, err := stmt.Exec(
			entry.ID,
			entry.Email,
			appstrings.RandomString(),
			"somesalt",
			time.Now(),
			time.Now(),
			dbservice.DbTrue,
		)
		if err != nil {
			log.Fatal(fmt.Sprintf("createTestUsers:2 %s", err.Error()))
		}
	}
}

func createTestAssets(db *dbservice.DB) {
	// @qusai what is maxWithdrawal for, can it be 0?,
	// it should habe a default value in the DB??
	stmt, err := db.Prepare(`INSERT INTO trade_asset
		SET
			id = ?,
			name = ?,
			symbol = ?,
			minDeposit = 0,
			maxWithdrawal = 0
		`,
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("createTestAssets:1 %s", err.Error()))
	}
	for _, entry := range testDbAssets {
		_, err := stmt.Exec(entry.ID, entry.Name, entry.Symbol)
		if err != nil {
			log.Fatal(fmt.Sprintf("createTestAssets:2 %s", err.Error()))
		}
	}
}
func createTestMarkets(db *dbservice.DB) {
	stmt, err := db.Prepare(
		`INSERT INTO trade_market SET
			id = ?,
			assetId = ?,
			baseAssetId = ?,
			isActive= ?`,
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("createTestMarkets:1 %s", err.Error()))
	}
	for _, entry := range testDbMarkets {
		_, err := stmt.Exec(
			entry.ID,
			entry.AssetID,
			entry.BaseAssetID,
			entry.IsActive,
		)
		if err != nil {
			log.Fatal(fmt.Sprintf("createTestMarkets:2 %s", err.Error()))
		}
	}
}

func truncateTable(db *dbservice.DB, table string) {
	stmt, err := db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE 1", table))
	if err == nil {
		_, err := stmt.Exec()
		if err != nil {
			log.Fatal("db_test:truncateTable:1 table: %s, error: %s", table, err.Error())
		}
	}
}

func dropTable(db *dbservice.DB, table string) {
	stmt, err := db.Prepare(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
	if err != nil {
		log.Fatal(fmt.Sprintf("db_test:dropTable:1 %s, %s", table, err.Error()))
	}
	{
		_, err := stmt.Exec()
		if err != nil {
			log.Fatal(fmt.Sprintf("db_test:dropTable:2 %s, %s", table, err.Error()))
		}
	}
}
