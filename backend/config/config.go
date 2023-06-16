package config

import (
	"github.com/caarlos0/env/v8"
	"log"
	"sync"
)

type Config struct {
	LambdaTaskRoot string `env:"LambdaTaskRoot" envDefault:""`
	ServiceAddress string `env:"ServiceAddress" envDefault:"0.0.0.0:4000"`
}

var initiated *Config

var once sync.Once

func read() *Config {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}

func GetInstance() *Config {
	once.Do(func() {
		initiated = read()
	})
	return initiated
}
