package models_test

import (
	"testing"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/FuturICT2/fin4-core/server/models"
)

func Test_AddOrder1(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1)) // BTC-USD
	user, _ := um.FindByID(models.ID(1))
	quantity := decimaldt.NewFromFloat(0.12345678)
	price := decimaldt.NewFromFloat(120.12345678)
	newOrder := models.NewOrder(market.ID, user.ID, models.BUY, quantity, price)

	{ // No market
		_, err := tm.AddOrder(nil, newOrder)
		if err != models.datatype.ErrMarketCantBeNil {
			t.Errorf("Test_AddOrder1 tm.AddOrder should return datatype.ErrMarketCantBeNil error")
		}
	}

	{ // No order
		_, err := tm.AddOrder(market, nil)
		if err != models.datatype.ErrOrderCantBeNil {
			t.Errorf("Test_AddOrder1 tm.AddOrder should return datatype.ErrOrderCantBeNil error")
		}
	}

	{ // No balance
		_, err := tm.AddOrder(market, newOrder)
		if err != models.datatype.ErrNotEnoughBalance {
			t.Errorf("Test_AddOrder1 tm.AddOrder should return datatype.ErrNotEnoughBalance error")
		}
	}

	{ // Should work
		um.DepositBalance(user.ID, market.BaseAssetID, decimaldt.NewFromFloat(1000))

		order, err := tm.AddOrder(market, newOrder)
		if err != nil {
			t.Errorf("Test_AddOrder1 tm.AddOrder should be successful, got %s", err.Error())
		}

		forder, err := tm.FindOrder(order.ID)
		if err != nil {
			t.Error("Test_AddOrder1 tm.FindOrder should not return an error")
		}
		if market.ID != forder.MarketID {
			t.Errorf("AddOrder -> invalid marketID, expected: %d, got: %d", market.ID, forder.MarketID)
		}
		if user.ID != forder.UserID {
			t.Errorf("AddOrder -> invalid userID, expected: %d, got: %d", user.ID, forder.UserID)
		}
		if models.BUY != forder.OrderType {
			t.Errorf("AddOrder -> invalid OrderType, expected: BUY")
		}
		if !forder.Quantity.Equals(quantity) {
			t.Errorf("AddOrder -> invalid Quantity, expected: %s, got: %s", quantity.String(), forder.Quantity.String())
		}
		if !forder.QuantityToFill.Equals(quantity) {
			t.Errorf("AddOrder -> invalid QuantityToFill, expected: %s, got: %s", quantity.String(), forder.QuantityToFill.String())
		}
		if !forder.Price.Equals(price) {
			t.Errorf("AddOrder -> invalid Price, expected: %s, got: %s", price.String(), forder.Price.String())
		}
		baseTotal := quantity.Mul(price)
		if baseTotal.String() != forder.BaseTotal.String() {
			t.Errorf("AddOrder -> invalid BaseTotal, expected: %s, got: %s", baseTotal.String(), forder.BaseTotal.String())
		}
		if forder.ReservedBalance.Equals(baseTotal) {
			t.Errorf("AddOrder -> invalid ReservedBalance, expected: %s, got: %s", baseTotal.String(), forder.ReservedBalance.String())
		}
		if !forder.ConsumedBalance.Equals(decimaldt.NewFromFloat(0.0)) {
			t.Errorf("AddOrder -> invalid ConsumedBalance, expected: 0, got: %d", forder.ConsumedBalance)
		}
	}
}

func Test_AddOrder2(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1)) // BTC-USD
	user, _ := um.FindByID(models.ID(1))
	quantity := decimaldt.NewFromFloat(2.12345678)
	price := decimaldt.NewFromFloat(20.12345678)
	newOrder := models.NewOrder(market.ID, user.ID, models.SELL, quantity, price)

	{ // No balance
		_, err := tm.AddOrder(market, newOrder)
		if err != models.datatype.ErrNotEnoughBalance {
			t.Errorf("Test_AddOrder2:AddOrder should return datatype.ErrNotEnoughBalance error")
		}
	}

	{ // Should work
		um.DepositBalance(user.ID, market.AssetID, decimaldt.NewFromFloat(1000))
		order, err := tm.AddOrder(market, newOrder)
		if err != nil {
			t.Errorf("Test_AddOrder2:AddOrder should be successful, got %s", err.Error())
		}

		forder, _ := tm.FindOrder(order.ID)
		if err != nil {
			t.Error("Test_AddOrder2:FindOrder should not return an error")
		}
		if models.SELL != forder.OrderType {
			t.Errorf("Test_AddOrder2:AddOrder -> invalid OrderType, expected: BUY")
		}
	}
}

func Test_AddCancelOrder_Balance(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1)) // BTC-USD
	user, _ := um.FindByID(models.ID(1))
	quantity := decimaldt.NewFromFloat(2.12345678)
	price := decimaldt.NewFromFloat(20.12345678)
	assetBalance := decimaldt.NewFromFloat(1000)
	baseAssetBalance := decimaldt.NewFromFloat(1000)
	um.DepositBalance(user.ID, market.AssetID, assetBalance)
	um.DepositBalance(user.ID, market.BaseAssetID, baseAssetBalance)
	{ // Buy acts on baseAssetBalance
		newOrder := models.NewOrder(market.ID, user.ID, models.BUY, quantity, price)
		order, _ := tm.AddOrder(market, newOrder)
		newBalance := baseAssetBalance.Sub(newOrder.BaseTotal)
		balance, _, _ := um.FindUserBalance(user.ID, market.BaseAssetID)
		if newBalance.String() != balance.String() {
			// Invalid new baseAssetBalance
			t.Errorf("Test_AddCancelOrder_Balance:1 %s ! %s", newBalance.String(), balance.String())
		}
		// assetBalance balance should not be affected
		balance, _, _ = um.FindUserBalance(user.ID, market.AssetID)
		if assetBalance.String() != balance.String() {
			// assetBalance should not be affected
			t.Errorf("Test_AddCancelOrder_Balance:2 %s ! %s", newBalance.String(), balance.String())
		}
		tm.CancelOrder(market, order)
		balance, _, _ = um.FindUserBalance(user.ID, market.BaseAssetID)
		if baseAssetBalance.String() != balance.String() {
			// cancel order should reset baseAssetBalance
			t.Errorf("Test_AddCancelOrder_Balance:3 %s ! %s", newBalance.String(), balance.String())
		}
	}
	{ // Sell acts on AssetBalance
		newOrder := models.NewOrder(market.ID, user.ID, models.SELL, quantity, price)
		order, _ := tm.AddOrder(market, newOrder)
		newBalance := assetBalance.Sub(newOrder.Quantity)
		balance, _, _ := um.FindUserBalance(user.ID, market.AssetID)
		if newBalance.String() != balance.String() {
			// Invalid new assetBalance
			t.Errorf("Test_AddCancelOrder_Balance:4 %s ! %s", newBalance.String(), balance.String())
		}
		// BaseAsset balance should not be affected
		balance, _, _ = um.FindUserBalance(user.ID, market.BaseAssetID)
		if assetBalance.String() != balance.String() {
			// baseAssetBalance should not be affected
			t.Errorf("Test_AddCancelOrder_Balance:5 %s ! %s", newBalance.String(), balance.String())
		}
		tm.CancelOrder(market, order)
		balance, _, _ = um.FindUserBalance(user.ID, market.AssetID)
		if baseAssetBalance.String() != balance.String() {
			// cancel order should reset assetBalance
			t.Errorf("Test_AddCancelOrder_Balance:6 %s ! %s", newBalance.String(), balance.String())
		}
	}
}

// Test_AddOrder_Overlap test that opposite overlapping orders are cancelled
func Test_AddOrder_Overlap(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1))
	user, _ := um.FindByID(models.ID(1))
	um.DepositBalance(user.ID, market.AssetID, decimaldt.NewFromFloat(1000))
	um.DepositBalance(user.ID, market.BaseAssetID, decimaldt.NewFromFloat(1000))
	quantity := decimaldt.NewFromFloat(1)
	{
		// user wants to sell at price X & he has an order to buy at price Y WHERE X <= Y
		// then it cancel buy order
		bOrder1, _ := tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.BUY, quantity, decimaldt.NewFromFloat(12)),
		)
		bOrder2, _ := tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.BUY, quantity, decimaldt.NewFromFloat(10)),
		)
		bOrder3, _ := tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.BUY, quantity, decimaldt.NewFromFloat(8)),
		)
		tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.SELL, quantity, decimaldt.NewFromFloat(9)),
		)
		if order, _ := tm.FindOrder(bOrder1.ID); order != nil {
			//overlaping buy order should be cancelled
			t.Error("Test_AddOrder_Overlap:1")
		}
		if order, _ := tm.FindOrder(bOrder2.ID); order != nil {
			// overlaping buy order should be cancelled
			t.Error("Test_AddOrder_Overlap:2")
		}
		if order, _ := tm.FindOrder(bOrder3.ID); order == nil {
			//none-overlapping buy order should NOT be cancelled
			t.Error("Test_AddOrder_Overlap:3")
		}
	}

	{
		// user wants to buy at price X & he has an order to sell at price Y WHERE X >= Y
		// then it cancel buy order
		sOrder1, _ := tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.SELL, quantity, decimaldt.NewFromFloat(8)),
		)
		sOrder2, _ := tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.SELL, quantity, decimaldt.NewFromFloat(10)),
		)
		sOrder3, _ := tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.SELL, quantity, decimaldt.NewFromFloat(12)),
		)
		tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.BUY, quantity, decimaldt.NewFromFloat(11)),
		)
		if order, _ := tm.FindOrder(sOrder1.ID); order != nil {
			//overlaping sell order should be cancelled
			t.Error("Test_AddOrder_Overlap:4")
		}
		if order, _ := tm.FindOrder(sOrder2.ID); order != nil {
			// overlaping sell order should be cancelled
			t.Error("Test_AddOrder_Overlap:5")
		}
		if order, _ := tm.FindOrder(sOrder3.ID); order == nil {
			//none-overlapping sell order should NOT be cancelled
			t.Error("Test_AddOrder_Overlap:6")
		}
	}

}

func Test_CancelOrder(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1))  // BTC-USD
	market2, _ := tm.FindMarket(models.ID(2)) // BTC-EUR
	user, _ := um.FindByID(models.ID(1))
	quantity := decimaldt.NewFromFloat(2.12345678)
	price := decimaldt.NewFromFloat(20.12345678)
	assetBalance := decimaldt.NewFromFloat(1000)
	baseAssetBalance := decimaldt.NewFromFloat(1000)
	um.DepositBalance(user.ID, market.AssetID, assetBalance)
	um.DepositBalance(user.ID, market.BaseAssetID, baseAssetBalance)
	order, _ := tm.AddOrder(market, models.NewOrder(market.ID, user.ID, models.SELL, quantity, price))

	// No Market
	if err := tm.CancelOrder(nil, order); err != models.datatype.ErrMarketCantBeNil {
		t.Error("Test_CancelOrder:1 expecting err to be datatype.ErrMarketCantBeNil")
	}

	// No Market
	if err := tm.CancelOrder(market, nil); err != models.datatype.ErrOrderCantBeNil {
		t.Error("Test_CancelOrder:2 expecting err to be datatype.ErrOrderCantBeNil")
	}

	// Wrong Market
	if err := tm.CancelOrder(market2, order); err != models.datatype.ErrOrderDeosntBelongToMarket {
		t.Error("Test_CancelOrder:3 expecting err to be datatype.ErrOrderDeosntBelongToMarket")
	}

	// should cancel
	if err := tm.CancelOrder(market, order); err != nil {
		t.Error("Test_CancelOrder:4 should cancel!")
	}

	// should be deleted
	if dbOrder, _ := tm.FindOrder(order.ID); dbOrder != nil {
		t.Error("Test_CancelOrder:5 should be deleted!")
	}
}

func Test_FindLowestSellOrder(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1))
	user, _ := um.FindByID(models.ID(1))
	um.DepositBalance(user.ID, market.AssetID, decimaldt.NewFromFloat(1000))
	um.DepositBalance(user.ID, market.BaseAssetID, decimaldt.NewFromFloat(1000))
	quantity := decimaldt.NewFromFloat(1)
	minPrice := 1.0
	minPriceDt := decimaldt.NewFromFloat(minPrice)
	maxPrice := 25.0
	maxPriceDt := decimaldt.NewFromFloat(maxPrice)

	for _, price := range []float64{5, maxPrice, minPrice, 15, 10} {
		tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.SELL, quantity, decimaldt.NewFromFloat(price)),
		)
	}
	order := tm.FindLowestSellOrder(market)
	if order == nil {
		// should find an lowest sell order
		t.Error("Test_FindLowestSellOrder:1")
	}
	if !order.Price.Equal(minPriceDt) {
		// should find correct lowest sell order
		t.Error("Test_FindLowestSellOrder:2")
	}

	if order := tm.FindLowestSellOrder(nil); order != nil {
		// market -> nil ? then order -> nil
		t.Error("Test_FindLowestSellOrder:3")
	}

	// FindLowestSellOrders
	{ // Market = nil ? orders -> nil
		orders := tm.FindLowestSellOrders(nil, 3)
		if len(orders) != 0 {
			t.Error("Test_FindLowestSellOrder:4 expected: %d orders, got: %d", 0, len(orders))
		}

	}
	{ // limit less than available
		orders := tm.FindLowestSellOrders(market, 3)
		if len(orders) != 3 {
			t.Error("Test_FindLowestSellOrder:5 expected: %d orders, got: %d", 3, len(orders))
		}
	}

	{ // limit more than available
		orders := tm.FindLowestSellOrders(market, 10)
		if len(orders) != 5 {
			t.Error("Test_FindLowestSellOrder:6 expected: %d orders, got: %d", 5, len(orders))
		}
	}

	{ // ordered??
		orders := tm.FindLowestSellOrders(market, 10)
		if !orders[0].Price.Equal(minPriceDt) {
			t.Error("Test_FindLowestSellOrder:7 expected: %s orders, got: %s", minPriceDt.String(), orders[0].Price.String())
		}
		if !orders[len(orders)-1].Price.Equal(maxPriceDt) {
			t.Error("Test_FindLowestSellOrder:8 expected: %s orders, got: %s", maxPriceDt.String(), orders[len(orders)-1].Price.String())
		}
	}
}

func Test_FindHighestBuyOrder(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1))
	user, _ := um.FindByID(models.ID(1))
	um.DepositBalance(user.ID, market.AssetID, decimaldt.NewFromFloat(1000))
	um.DepositBalance(user.ID, market.BaseAssetID, decimaldt.NewFromFloat(1000))
	quantity := decimaldt.NewFromFloat(1)
	minPrice := 1.0
	minPriceDt := decimaldt.NewFromFloat(minPrice)
	maxPrice := 25.0
	maxPriceDt := decimaldt.NewFromFloat(maxPrice)
	for _, price := range []float64{5, maxPrice, minPrice, 15, 10} {
		tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.BUY, quantity, decimaldt.NewFromFloat(price)),
		)
	}

	// FindHighestBuyOrder
	order := tm.FindHighestBuyOrder(market)
	if order == nil {
		// should find an highest buy order
		t.Error("Test_FindHighestBuyOrder:1")
	}
	if !order.Price.Equal(decimaldt.NewFromFloat(25)) {
		// should find correct highest buy order
		t.Error("Test_FindHighestBuyOrder:2")
	}
	if order := tm.FindHighestBuyOrder(nil); order != nil {
		// market -> nil ? then order -> nil
		t.Error("Test_FindHighestBuyOrder:3")
	}

	// FindHighestBuyOrders
	{ // Market = nil ? orders -> nil
		orders := tm.FindHighestBuyOrders(nil, 3)
		if len(orders) != 0 {
			t.Error("Test_FindHighestBuyOrder:4 expected: %d orders, got: %d", 0, len(orders))
		}

	}
	{ // limit less than available
		orders := tm.FindHighestBuyOrders(market, 3)
		if len(orders) != 3 {
			t.Error("Test_FindHighestBuyOrder:5 expected: %d orders, got: %d", 3, len(orders))
		}
	}

	{ // limit more than available
		orders := tm.FindHighestBuyOrders(market, 10)
		if len(orders) != 5 {
			t.Error("Test_FindHighestBuyOrder:6 expected: %d orders, got: %d", 5, len(orders))
		}
	}

	{ // ordered??
		orders := tm.FindHighestBuyOrders(market, 10)
		if !orders[0].Price.Equal(maxPriceDt) {
			t.Error("Test_FindHighestBuyOrder:7 expected: %s orders, got: %s", maxPriceDt.String(), orders[0].Price.String())
		}
		if !orders[len(orders)-1].Price.Equal(minPriceDt) {
			t.Error("Test_FindHighestBuyOrder:8 expected: %s orders, got: %s", minPriceDt.String(), orders[len(orders)-1].Price.String())
		}
	}
}

func Test_FindLowerBuyOrders(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1))
	user, _ := um.FindByID(models.ID(1))
	um.DepositBalance(user.ID, market.AssetID, decimaldt.NewFromFloat(1000))
	um.DepositBalance(user.ID, market.BaseAssetID, decimaldt.NewFromFloat(1000))
	quantity := decimaldt.NewFromFloat(1)

	for _, price := range []float64{5, 25, 20, 15, 10, 30, 35} {
		tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.BUY, quantity, decimaldt.NewFromFloat(price)),
		)
	}
	order, _ := tm.AddOrder(
		market,
		models.NewOrder(market.ID, user.ID, models.BUY, quantity, decimaldt.NewFromFloat(30)),
	)

	// market nil -> orders []
	if orders := tm.FindLowerBuyOrders(nil, int64(order.ID), 3); orders != nil {
		t.Error("Test_FindLowerBuyOrders:1")
	}
	orders := tm.FindLowerBuyOrders(market, int64(order.ID), 3)
	if len(orders) != 3 {
		t.Error("Test_FindLowerBuyOrders:2 len(orders) = %d", len(orders))
	}
	if !orders[0].Price.Equals(decimaldt.NewFromFloat(25)) {
		t.Error("Test_FindLowerBuyOrders:4 first order price should be 25")
	}
	if !orders[1].Price.Equals(decimaldt.NewFromFloat(20)) {
		t.Error("Test_FindLowerBuyOrders:5 second order price should be 20")
	}
	if !orders[2].Price.Equals(decimaldt.NewFromFloat(15)) {
		t.Error("Test_FindLowerBuyOrders:6 third order price should be 15")
	}
}

func Test_FindHigherSellOrders(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1))
	user, _ := um.FindByID(models.ID(1))
	um.DepositBalance(user.ID, market.AssetID, decimaldt.NewFromFloat(1000))
	um.DepositBalance(user.ID, market.BaseAssetID, decimaldt.NewFromFloat(1000))
	quantity := decimaldt.NewFromFloat(1)

	for _, price := range []float64{5, 25, 20, 15, 10, 30, 35} {
		tm.AddOrder(
			market,
			models.NewOrder(market.ID, user.ID, models.SELL, quantity, decimaldt.NewFromFloat(price)),
		)
	}
	order, _ := tm.AddOrder(
		market,
		models.NewOrder(market.ID, user.ID, models.SELL, quantity, decimaldt.NewFromFloat(16)),
	)

	// market nil -> orders []
	if orders := tm.FindHigherSellOrders(nil, int64(order.ID), 3); orders != nil {
		t.Error("Test_FindHigherSellOrders:1")
	}
	orders := tm.FindHigherSellOrders(market, int64(order.ID), 3)
	if len(orders) != 3 {
		t.Error("Test_FindHigherSellOrders:2 len(orders) = %d", len(orders))
	}
	if !orders[0].Price.Equals(decimaldt.NewFromFloat(20)) {
		t.Error("Test_FindHigherSellOrders:4 first order price should be 25")
	}
	if !orders[1].Price.Equals(decimaldt.NewFromFloat(25)) {
		t.Error("Test_FindHigherSellOrders:5 second order price should be 20")
	}
	if !orders[2].Price.Equals(decimaldt.NewFromFloat(30)) {
		t.Error("Test_FindHigherSellOrders:6 third order price should be 15")
	}
}

func Test_FindUserOrders(t *testing.T) {
	db := SetupTestDB()
	defer TearDownTestDB(db)
	tm := db.NewTradeModel()
	um := db.NewUserModel()
	market, _ := tm.FindMarket(models.ID(1))
	u1ID := models.ID(1)
	u2ID := models.ID(2)
	um.DepositBalance(u1ID, market.AssetID, decimaldt.NewFromFloat(1000))
	um.DepositBalance(u1ID, market.BaseAssetID, decimaldt.NewFromFloat(1000))
	um.DepositBalance(u2ID, market.AssetID, decimaldt.NewFromFloat(1000))
	um.DepositBalance(u2ID, market.BaseAssetID, decimaldt.NewFromFloat(1000))
	quantity := decimaldt.NewFromFloat(1)
	var orderType models.OrderType
	for idx, price := range []float64{1, 2, 3, 4, 5} {
		if idx > 3 {
			orderType = models.SELL
		} else {
			orderType = models.BUY
		}
		tm.AddOrder(
			market,
			models.NewOrder(market.ID, u1ID, orderType, quantity, decimaldt.NewFromFloat(price)),
		)
	}
	for idx, price := range []float64{1, 2, 3, 4, 5} {
		if idx > 3 {
			orderType = models.SELL
		} else {
			orderType = models.BUY
		}
		tm.AddOrder(
			market,
			models.NewOrder(market.ID, u2ID, orderType, quantity, decimaldt.NewFromFloat(price)),
		)
	}
	{
		allCount, orders, err := tm.FindUserOrders(u1ID, 0, 0, 100)
		if err != nil {
			t.Error("Test_FindUserOrders:1 should not return error")
		}
		if allCount != 5 {
			t.Error("Test_FindUserOrders:2 allCount ! %d", allCount)
		}
		if len(orders) != 5 {
			t.Error("Test_FindUserOrders:3 len(order) ! %d", len(orders))
		}
		for _, entry := range orders {
			if entry.UserID != u1ID {
				t.Error("Test_FindUserOrders:4 got: %d", int64(u1ID))
			}
		}
	}

	{
		allCount, orders, err := tm.FindUserOrders(u1ID, 0, 0, 2)
		if err != nil {
			t.Error("Test_FindUserOrders:5 should not return error")
		}
		if allCount != 5 {
			t.Error("Test_FindUserOrders:6 allCount ! %d", allCount)
		}
		if len(orders) != 2 {
			t.Error("Test_FindUserOrders:7 len(order) ! %d", len(orders))
		}
	}
}
