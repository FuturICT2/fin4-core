package env

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lytics/logrus"
)

var envLoaded = false

//Load load .env file
func Load(baseDir string, isTesting bool) {
	var envFile string
	if isTesting {
		envFile = "env-test"
	} else {
		envFile = "env-default"
	}

	err := godotenv.Load(os.ExpandEnv(
		fmt.Sprintf("%s/%s", baseDir, envFile),
	))
	if err != nil {
		logrus.Warn(fmt.Sprintf("Failed loading env file, error: %d", err.Error()))
	}
	envLoaded = true
	//checkRequiredEnvs()
}

// Getenv Gets env var value
func Getenv(varName string) string {
	return os.Getenv(varName)
}

// MustGetenv Gets env var value, will panic if var is not set
func MustGetenv(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		panic(fmt.Sprintf("Can't get ENV variable `%s`", varName))
	}
	return value
}

func checkRequiredEnvs() {
	forceEnv := strings.TrimSpace(Getenv("FORCE_ENV_CHECKING"))
	if forceEnv != "1" && forceEnv != "0" {
		log.Fatal("env var FORCE_ENV_CHECKING is missing! Set it to one of 1|0")
		return
	}
	if forceEnv == "0" {
		return
	}
	requiredVars := []string{
		"BASE_URL",
		"EMAIL_FROM",
		"AWS_SES_KEY",
		"AWS_SES_SECRET",
		"SLACK_WEBHOOK_URL",
	}
	missingFatal := []string{}
	for _, val := range requiredVars {
		if Getenv(val) == "" {
			missingFatal = append(missingFatal, val)
		}
	}
	if len(missingFatal) > 0 {
		log.Fatal(
			fmt.Sprintf("Some required env vars are missing [%s]", strings.Join(missingFatal, ", ")),
		)
	}
}
