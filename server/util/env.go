package util

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var envLoaded = false

func init() {
	checkRequiredEnvs()
}

// LoadEnv load .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: loading .env file", err.Error())
	} else {
		envLoaded = true
	}
}

// Getenv Gets env var value
func Getenv(varName string) string {
	if !envLoaded {
		LoadEnv()
	}
	return os.Getenv(varName)
}

// MustGetenv Gets env var value, will panic if var is not set
func MustGetenv(varName string) string {
	value := Getenv(varName)
	if value == "" {
		panic(fmt.Sprintf("Can't get ENV variable `%s`", varName))
	}
	return value
}

func checkRequiredEnvs() {
	forceEnv := strings.TrimSpace(Getenv("FORCE_ENV_CHECKING"))
	log.Println("Force environment variables checking = ", forceEnv)
	if forceEnv != "1" && forceEnv != "0" {
		log.Fatal("env var FORCE_ENV_CHECKING is missing! Set it to one of 1|0")
		return
	}
	if forceEnv == "0" {
		return
	}
	requiredVars := []string{
		"BASE_URL",
		"ETHEREUM_NET",
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
