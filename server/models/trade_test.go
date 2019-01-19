package models_test

import (
	"testing"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/FuturICT2/fin4-core/server/models"
)

func Test_OrderType(t *testing.T) {
	if models.BUY == models.SELL {
		t.Error("BUY value should not equals SELL value")
	}
}

func Test_NewOrder(t *testing.T) {
	marketID := models.ID(1)
	userID := models.ID(2)
	quantity := decimaldt.NewFromFloat(0.12345678)
	price := decimaldt.NewFromFloat(1.12345678)
	o := models.NewOrder(marketID, userID, models.BUY, quantity, price)
	if o.MarketID != marketID {
		t.Error("NewOrder should set MarketID correctly")
	}
	if o.UserID != userID {
		t.Error("NewOrder should set userID correctly")
	}
	if !o.Quantity.Equals(quantity) {
		t.Error("NewOrder should set Quantity correctly")
	}
	if !o.QuantityToFill.Equals(quantity) {
		t.Error("NewOrder should set QuantityToFill correctly")
	}

	if !o.Price.Equals(price) {
		t.Error("NewOrder should set userID correctly")
	}
	if o.OrderType != models.BUY {
		t.Error("NewOrder -> OrderType should be BUY")
	}
	o = models.NewOrder(marketID, userID, models.SELL, quantity, price)
	if o.OrderType != models.SELL {
		t.Error("NewOrder -> OrderType should be SELL")
	}
}
