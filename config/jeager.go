package config

import "os"

func initJaeger(app *AppConfig) {
	host := os.Getenv("JAEGER_HOST")
	port := os.Getenv("JAEGER_PORT")

	app.Jaeger.Host = host
	app.Jaeger.Port = port
}
