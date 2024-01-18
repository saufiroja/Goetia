package config

import "os"

func initHttp(conf *AppConfig) {
	port := os.Getenv("HTTP_PORT")

	conf.Http.Port = port
}
