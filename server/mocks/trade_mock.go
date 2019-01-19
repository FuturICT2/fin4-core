package models

import "github.com/FuturICT2/fin4-core/server/datatype"

// TradeModelMock TradeModelMock datatype
type TradeModelMock struct {
	FindAssetsResult []datatype.Asset
	FindAssetsError  error

	FindUserFavoriteAssetsResult []datatype.Asset
	FindUserFavoriteAssetsError  error

	ToggleFavoriteAssetErr  error
	GetAssetAveragePriceErr error
	FindAssetResult         *datatype.Asset
	FindAssetError          error

	FindAssetBySymbolResult *datatype.Asset
	FindAssetBySymbolError  error

	FindMarketResult *datatype.Market
	FindMarketError  error

	FindMarketsResult []datatype.Market
	FindMarketsError  error

	AddOrderResult *datatype.Order
	AddOrderError  error

	CancelOrderError error

	FindOrderResult *datatype.Order
	FindOrderError  error

	FindLowestSellOrderResult *datatype.Order
	FindHighestBuyOrderResult *datatype.Order

	FindLowerBuyOrdersResult []*datatype.Order

	FindHigherSellOrdersResult []*datatype.Order

	FindLowestSellOrdersResult []*datatype.Order
	FindHighestBuyOrdersResult []*datatype.Order

	FindUserOrdersAllCount int
	FindUserOrdersResult   []*datatype.Order
	FindUserOrdersError    error

	FindUserTransactionsTotal  int
	FindUserTransactionsResult []datatype.Transaction
	FindUserTransactionsError  error

	FindLatestTransactionsResult []datatype.Transaction
	FindLatestTransactionsError  error
}

// FindAssets FindAssets mock
func (ts *TradeModelMock) FindAssets() ([]datatype.Asset, error) {
	return ts.FindAssetsResult, ts.FindAssetsError
}

func (ts *TradeModelMock) ToggleFavoriteAsset(*datatype.User, datatype.ID) error {
	return ts.ToggleFavoriteAssetErr
}

//FindUserFavoriteAssets mock
func (ts *TradeModelMock) FindUserFavoriteAssets(user *datatype.User) ([]datatype.Asset, error) {
	return ts.FindUserFavoriteAssetsResult, ts.FindUserFavoriteAssetsError
}

// FindAsset FindAsset mock
func (ts *TradeModelMock) FindAsset(id datatype.ID) (*datatype.Asset, error) {
	return ts.FindAssetResult, ts.FindAssetError
}

// InsertAsset InsertAsset mock
func (ts *TradeModelMock) InsertAsset(
	userID datatype.ID,
	name string,
	symbol string,
	description string,
	shares string,
	category string,
	hashtags []string,
	logoFile string,
) (*datatype.Market, error) {
	return nil, nil
}

// FindAssetBySymbol FindAsset mock
func (ts *TradeModelMock) FindAssetBySymbol(symbol string) (*datatype.Asset, error) {
	return ts.FindAssetBySymbolResult, ts.FindAssetBySymbolError
}

// FindMarkets FindMarkets mock
func (ts *TradeModelMock) FindMarkets() ([]datatype.Market, error) {
	return ts.FindMarketsResult, ts.FindMarketsError
}

// FindMarket FindMarket mock
func (ts *TradeModelMock) FindMarket(datatype.ID) (*datatype.Market, error) {
	return ts.FindMarketResult, ts.FindMarketError
}

// FindMarketByAssetID FindMarket mock
func (ts *TradeModelMock) FindMarketByAssetID(datatype.ID) (*datatype.Market, error) {
	return ts.FindMarketResult, ts.FindMarketError
}

// AddOrder AddOrder mock
func (ts *TradeModelMock) AddOrder(datatype.UserService, *datatype.Market, *datatype.Order) (*datatype.Order, error) {
	return ts.AddOrderResult, ts.AddOrderError
}

// CancelOrder CancelOrder mock
func (ts *TradeModelMock) CancelOrder(*datatype.Market, *datatype.Order) error {
	return ts.CancelOrderError
}

// FindOrder FindOrder mock
func (ts *TradeModelMock) FindOrder(datatype.ID) (*datatype.Order, error) {
	return ts.FindOrderResult, ts.FindOrderError
}

// ExecuteTransaction ExecuteTransaction mock
func (ts *TradeModelMock) ExecuteTransaction(userService datatype.UserService, m *datatype.Market, buy *datatype.Order, sell *datatype.Order) bool {
	return true
}

// FindLowestSellOrder FindLowestSellOrder mock
func (ts *TradeModelMock) FindLowestSellOrder(*datatype.Market) *datatype.Order {
	return ts.FindLowestSellOrderResult
}

// FindHighestBuyOrder FindHighestBuyOrder mock
func (ts *TradeModelMock) FindHighestBuyOrder(*datatype.Market) *datatype.Order {
	return ts.FindHighestBuyOrderResult
}

// FindLowestSellOrders FindLowestSellOrders mock
func (ts *TradeModelMock) FindLowestSellOrders(*datatype.Market, int) []*datatype.Order {
	return ts.FindLowestSellOrdersResult
}

// FindHighestBuyOrders FindHighestBuyOrders mock
func (ts *TradeModelMock) FindHighestBuyOrders(*datatype.Market, int) []*datatype.Order {
	return ts.FindHighestBuyOrdersResult
}

// FindLowerBuyOrders FindLowerBuyOrders mock
func (ts *TradeModelMock) FindLowerBuyOrders(*datatype.Market, int64, int) []*datatype.Order {
	return ts.FindLowerBuyOrdersResult
}

// FindHigherSellOrders FindHigherSellOrders mock
func (ts *TradeModelMock) FindHigherSellOrders(*datatype.Market, int64, int) []*datatype.Order {
	return ts.FindHigherSellOrdersResult
}

// FindUserOrders  mock
func (ts *TradeModelMock) FindUserOrders(
	userID datatype.ID,
	marketID datatype.ID,
	offset int,
	limit int,
) (int, []*datatype.Order, error) {
	return ts.FindUserOrdersAllCount, ts.FindUserOrdersResult, ts.FindUserOrdersError
}

// FindMarketOrders  mock
func (ts *TradeModelMock) FindMarketOrders(
	marketID datatype.ID,
	offset int,
	limit int,
) (int, []*datatype.Order, error) {
	return ts.FindUserOrdersAllCount, ts.FindUserOrdersResult, ts.FindUserOrdersError
}

//FindLatestTransactions mock
func (ts *TradeModelMock) FindLatestTransactions(m *datatype.Market, lastTr *datatype.Transaction, offset int, limit int) ([]datatype.Transaction, error) {
	return ts.FindLatestTransactionsResult, ts.FindLatestTransactionsError
}

// FindUserTransactions  mock
func (ts *TradeModelMock) FindUserTransactions(userID datatype.ID, offset int, limit int) (int, []datatype.Transaction, error) {
	return ts.FindUserTransactionsTotal, ts.FindUserTransactionsResult, ts.FindUserTransactionsError
}

// MatchOrders match orders mock
func (ts *TradeModelMock) MatchOrders(datatype.UserService, *datatype.Market) (*datatype.Order, *datatype.Order) {
	return nil, nil
}

// UpdateMarketSummary mock
func (db *TradeModelMock) UpdateMarketSummary(*datatype.Market) error {
	return nil
}
