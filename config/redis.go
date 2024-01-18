package config

import "os"

func initRedis(conf *AppConfig) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASS")

	conf.Redis.Host = host
	conf.Redis.Port = port
	conf.Redis.Pass = pass
}
