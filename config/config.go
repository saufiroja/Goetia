package config

import "github.com/joho/godotenv"

type AppConfig struct {
	App struct {
		Env         string
		ServiceName string
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
	Jaeger struct {
		Host string
		Port string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	// add config file path in .env
	_ = godotenv.Load("../.env")

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initGrpc(appConfig)
		initHttp(appConfig)
		initPostgres(appConfig)
		initRedis(appConfig)
		initJaeger(appConfig)
	}

	return appConfig
}
