package main

import (
	"strings"

	"github.com/FuturICT2/fin4-core/server/assetservice"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/dbservice"
	"github.com/FuturICT2/fin4-core/server/emailer"
	"github.com/FuturICT2/fin4-core/server/env"
	"github.com/FuturICT2/fin4-core/server/ethereum"
	"github.com/FuturICT2/fin4-core/server/filestorage"
	"github.com/FuturICT2/fin4-core/server/logger"
	"github.com/FuturICT2/fin4-core/server/routes"
	"github.com/FuturICT2/fin4-core/server/timelineservice"
	"github.com/FuturICT2/fin4-core/server/userservice"
)

func main() {
	environment := strings.ToLower(env.MustGetenv("ENVIRONMENT"))

	//TODO this should be changed ASAP
	baseDir := "$GOPATH/fin4-core"
	env.Load(baseDir, false)

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

	// connect to database
	db := dbservice.MustConnect(cfg.DataSourceName)
	userService := userservice.NewService(db)
	timelineService := timelineservice.NewService(db)
	assetService := assetservice.NewService(db)
	// connect with Amazon AWS (used in our live instance)
	fs := filestorage.GetStorage(
		env.Getenv("AWS_SES_KEY"),
		env.Getenv("AWS_SES_SECRET"),
	)
	// connect with an instance of ethereum
	ethereum := ethereum.MustNewEthereum()
	sc := datatype.ServiceContainer{
		Config:          cfg,
		AssetService:    assetService,
		TimelineService: timelineService,
		UserService:     userService,
		FileStorage:     fs,
		Ethereum:        ethereum,
	}
	// start listening to api calls
	routes.SetupRouting(sc).Run(":" + cfg.ServerPort)
}

func getPort() string {
	port := env.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
