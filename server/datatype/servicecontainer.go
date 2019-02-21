package datatype

import "github.com/FuturICT2/fin4-core/server/ethereum"

//ServiceContainer defines our service container type
type ServiceContainer struct {
	Config          Config
	AssetService    AssetService
	TimelineService TimelineService
	UserService     UserService
	FileStorage     FileStorage
	Ethereum        *ethereum.Ethereum
}
