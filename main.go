package main

import (
	"os"

	"github.com/FuturICT2/fin4-core/server/ethereum"
	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/FuturICT2/fin4-core/server/pkg/filestorage"
	"github.com/FuturICT2/fin4-core/server/routes"
	"github.com/FuturICT2/fin4-core/server/util"
	"github.com/sirupsen/logrus"
)

func initLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	initLogger()
	db := models.MustConnect(util.MustGetenv("DATA_SOURCE_NAME"))
	ethereum := ethereum.MustNewEthereum()
	fs := filestorage.GetStorage(
		util.Getenv("AWS_SES_KEY"),
		util.Getenv("AWS_SES_SECRET"),
	)
	routesEnv := routes.Env{
		DB:          db,
		Ethereum:    ethereum,
		FileStorage: fs,
	}
	routesEnv.StartRouter().Run(":" + getPort())
}

func getPort() string {
	port := util.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
