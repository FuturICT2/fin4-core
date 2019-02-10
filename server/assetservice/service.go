package assetservice

import "github.com/FuturICT2/fin4-core/server/dbservice"

//Service defines asset service type
type Service struct {
	*dbservice.DB
}

//NewService factory for Service
func NewService(db *dbservice.DB) *Service {
	return &Service{db}
}

func getAssetCols() string {
	return `
		trade_asset.id,
		trade_asset.category,
		trade_asset.symbol,
		trade_asset.name,
		trade_asset.creatorId,
		trade_asset.description,
		trade_asset.totalSupply,
		trade_asset.marketCap,
		trade_asset.minConf,
		trade_asset.decimals,
		trade_asset.blockchainExplorerUrl,
		trade_asset.networkFee,
		trade_asset.isListed,
		trade_asset.isDepositEnabled,
		trade_asset.isWithdrawalEnabled,
		trade_asset.minersCounter`
}
