package models_test

import (
	"testing"

	"github.com/FuturICT2/fin4-core/server/models"
)

func Test_FindMarket(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	if _, err := tm.FindMarket(models.ID(99999)); err == nil {
		t.Error("Test_FindMarket:FindMarket should fail when  given an invalid ID")
	}

	m, err := tm.FindMarket(models.ID(testDbMarkets[0].ID))
	if err != nil {
		t.Errorf("Test_FindMarket:FindMarket should not return an error, Got %s", err.Error())
	}

	if "Bitcoin" != m.Asset || models.ID(1) != m.AssetID {
		t.Errorf(
			"Test_FindMarket:FindMarket should return correct market data, Expected Asset: 1 -> BITCOIN, Got: %d -> %s",
			m.AssetID,
			m.Asset,
		)
	}
	if "USD" != m.BaseAsset || models.ID(3) != m.BaseAssetID {
		t.Errorf(
			"Test_FindMarket:FindMarket should return correct market data, Expected BaseAsset: 3 -> USD, Got: %d -> %s",
			m.BaseAssetID,
			m.BaseAsset,
		)
	}
}

func Test_FindMarkets(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	markets, err := tm.FindMarkets()
	if err != nil {
		t.Error("FindMarkets should not return an error")
	}
	if len(markets) == 0 {
		t.Error("FindMarkets should not return empty set")
	}

	for _, entry := range markets {
		market, err := tm.FindMarket(entry.ID)
		if err != nil || market == nil || market.ID != entry.ID {
			t.Error("FindMarket not working")
		}
	}
}
