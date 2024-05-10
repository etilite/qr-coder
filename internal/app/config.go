package app

import "os"

type Config struct {
	HTTPAddr string
}

func NewConfigFromEnv() Config {
	config := Config{
		HTTPAddr: ":8080",
	}
	httpAddr, ok := os.LookupEnv("HTTP_ADDR")
	if ok {
		config.HTTPAddr = httpAddr
	}
	return config
}
