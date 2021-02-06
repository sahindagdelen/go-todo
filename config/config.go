package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

var Config Configuration

const configurationFilepath = "config.yaml"

type Configuration struct {
	Database struct {
		Host   string `yaml:"host" env:"DB_HOST" env-description:"Database host name"`
		Username   string `yaml:"username" env:"DB_USER" env-description:"Database user name"`
		Password   string `yaml:"password"  env:"DB_PASSWORD" env-description:"Database user password"`
		Name       string `yaml:"name" env:"DB_NAME" env-description:"Database name"`
		Collection string `yaml:"collection" env:"DB_NAME" env-description:"Collection name"`
	} `yaml:"database"`
}

func Initialize() {
	err := cleanenv.ReadConfig(configurationFilepath, &Config)
	if err != nil {
		panic(fmt.Sprintf("Failed to read configuration '%s', cannot continue - %s", configurationFilepath, err))
	}
}
