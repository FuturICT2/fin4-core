package models_test

import (
	"testing"

	"github.com/FuturICT2/fin4-core/server/models"
)

func Test_FindAssets(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	res, err := db.NewTradeModel().FindAssets()
	if err != nil {
		t.Error("FindAssets should not return an error", err)
	}
	if len(res) == 0 {
		t.Error("FindAssets should return some")
	}

}

func Test_FindAsset(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	asset, err := db.NewTradeModel().FindAsset(999999)
	if err == nil {
		t.Error("FindAsset should return an error on invalidID")
	}

	asset, err = db.NewTradeModel().FindAsset(models.ID(testDbAssets[0].ID))
	if asset == nil {
		t.Errorf("FindAsset should be successful on valid ID, ID: %d", models.ID(testDbAssets[0].ID))
	}
	if err != nil {
		t.Error("FindAsset should not return an error on success")
	}
	if asset.Name != testDbAssets[0].Name {
		t.Errorf(
			"FindAsset should return correct asset, name is wrong, expected: %s, got: %s",
			testDbAssets[0].Name,
			asset.Name,
		)
	}
	if asset.Symbol != testDbAssets[0].Symbol {
		t.Errorf(
			"FindAsset should return correct asset, name is wrong, expected: %s, got: %s",
			testDbAssets[0].Name,
			asset.Name,
		)
	}
}
