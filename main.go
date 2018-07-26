package main

import (
	"os"

	"github.com/FuturICT2/fin4-core/server/ethereum"
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
	ethereum := ethereum.MustNewEthereum()
	routesEnv := routes.Env{
		// We can add general types like database connection to the ENV when needed
		Ethereum: ethereum,
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
