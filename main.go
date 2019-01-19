package main

import (
	"strings"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/dbservice"
	"github.com/FuturICT2/fin4-core/server/emailer"
	"github.com/FuturICT2/fin4-core/server/env"
	"github.com/FuturICT2/fin4-core/server/ethereum"
	"github.com/FuturICT2/fin4-core/server/filestorage"
	"github.com/FuturICT2/fin4-core/server/logger"
	"github.com/FuturICT2/fin4-core/server/routes"
	"github.com/FuturICT2/fin4-core/server/tokenservice"
	"github.com/FuturICT2/fin4-core/server/userservice"
)

func main() {
	environment := strings.ToLower(env.MustGetenv("ENVIRONMENT"))

	//TODO this should be changed ASAP
	baseDir := "$GOPATH/src/github.com/FuturICT2/fin4-core"
	env.Load(baseDir, environment == "test")

	cfg := datatype.Config{
		Environment:        environment,
		DBMigrationPath:    env.MustGetenv("DB_MIGRATIONS_PATH"),
		ServerPort:         getPort(),
		DataSourceName:     env.MustGetenv("DATA_SOURCE_NAME"),
		TestDataSourceName: env.Getenv("TEST_DATA_SOURCE_NAME"),
		AwsSesKey:          env.Getenv("AWS_SES_KEY"),
		AwsSesSecret:       env.Getenv("AWS_SES_SECRET"),
		AwsEmailFrom:       env.Getenv("EMAIL_FROM"),
		SlackWebhookURL:    env.Getenv("SLACK_WEBHOOK_URL"),
	}

	logger.Setup(cfg)

	emailer.Setup(cfg)

	db := dbservice.MustConnect(cfg.DataSourceName)
	userService := userservice.NewService(db)
	tokenService := tokenservice.NewService(db)

	fs := filestorage.GetStorage(
		env.Getenv("AWS_SES_KEY"),
		env.Getenv("AWS_SES_SECRET"),
	)
	ethereum := ethereum.MustNewEthereum()
	sc := datatype.ServiceContainer{
		Config:       cfg,
		UserService:  userService,
		TokenService: tokenService,
		FileStorage:  fs,
		Ethereum:     ethereum,
	}

	routes.SetupRouting(sc).Run(":" + cfg.ServerPort)
}

func getPort() string {
	port := env.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
