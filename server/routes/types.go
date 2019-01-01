package routes

import (
	"github.com/FuturICT2/fin4-core/server/ethereum"
	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/FuturICT2/fin4-core/server/pkg/filestorage"
)

// Env router env type
type Env struct {
	Ethereum    *ethereum.Ethereum
	DB          *models.DB
	FileStorage *filestorage.Storage
}
