package datatype

// AssetService defines asset service interface
type AssetService interface {
	FindAll() ([]Asset, error)
	FindByID(id ID) (*Asset, error)
	FindBySymbol(symbol string) (*Asset, error)
	FindByName(name string) (*Asset, error)
	Create(
		sc ServiceContainer,
		userID ID,
		name string,
		symbol string,
		description string,
	) (*Asset, error)
	ToggleFavorite(user *User, assetID ID) error
	ToggleFavoriteBlock(user *User, blockID ID) error
	FindUserFavoriteAssets(user *User) ([]Asset, error)
	DidUserLikeBlock(userID ID, blockID ID) bool
	CreateAssetBlock(
		userID ID,
		assetID ID,
		blockText string,
		images []string,
	) (*Block, error)
	GetAssetBlocks(assetID ID) ([]Block, error)
	VerifyAssetBlock(sc *ServiceContainer, user *User, blockID ID, status int) error
	GetAssetMiners(assetID ID) ([]Miner, error)
	IsOracle(userID ID, assetID ID) bool
	GetAssetBlockImages(blockID ID) ([]string, error)
}
