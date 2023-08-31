package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var alConfigEnvName = "AL_CONFIG"

type config struct {
	Urls  map[string]string `yaml:"urls"`
	Open  map[string]string `yaml:"apps-to-open"`
	Close map[string]string `yaml:"apps-to-close"`
}

func ReadConfig() (*config, error) {
	var (
		cfg config
		err error
	)

	configPath := configPath(alConfigEnvName)

	if configNotExists(configPath) {
		log.Fatal(fmt.Printf("Config file %s does no exist ", configPath))
	}

	err = cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func configPath(envName string) string {
	if envName == "" {
		envName = alConfigEnvName
	}

	configPath := os.Getenv(envName)
	if configPath == "" {
		log.Fatal("Path is empty. Please specify config path in " + envName + " env variable")
	}

	return configPath
}

func configNotExists(configPath string) bool {
	_, err := os.Stat(configPath)
	return os.IsNotExist(err)
}
