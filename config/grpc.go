package config

import "os"

func initGrpc(conf *AppConfig) {
	port := os.Getenv("GRPC_PORT")

	conf.Grpc.Port = port
}
