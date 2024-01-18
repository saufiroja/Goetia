package config

import "github.com/joho/godotenv"

type AppConfig struct {
	App struct {
		Env string
	}
	Grpc struct {
		Port string
	}
	Http struct {
		Port string
	}
	Postgres struct {
		Name string
		User string
		Pass string
		Host string
		Port string
		SSL  string
	}
	Redis struct {
		Host string
		Port string
		Pass string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	// add config file path in .env
	err := godotenv.Load("../.env")
	if err != nil {
		panic("error loading .env file")
	}

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initGrpc(appConfig)
		initHttp(appConfig)
		initPostgres(appConfig)
		initRedis(appConfig)
	}

	return appConfig
}
