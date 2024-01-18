package config

import (
	"log"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func initApp(conf *AppConfig) {
	env := os.Getenv("GO_ENV")

	switch cases.Lower(language.English).String(env) {
	case "development":
		conf.App.Env = "development"
		log.Println("App environment is set to development")
	case "staging":
		conf.App.Env = "staging"
		log.Println("App environment is set to staging")
	case "testing":
		conf.App.Env = "testing"
		log.Println("App environment is set to testing")
	case "production":
		conf.App.Env = "production"
		log.Println("App environment is set to production")
	default:
		conf.App.Env = "development"
		log.Println("App environment is not set. Using default environment development")
	}

	conf.App.Env = env
}
