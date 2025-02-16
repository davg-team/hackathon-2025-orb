package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var cfg config

type config struct {
	EnvType   string        `yaml:"env_type"`
	PathToKey string        `yaml:"path_to_key"`
	Server    serverConfig  `yaml:"server"`
	Storage   storageConfig `yaml:"storage"`
}

type serverConfig struct {
	Port int `yaml:"port"`
}

// type storageConfig struct {
// 	Host   string `yaml:"host"`
// 	Port   int    `yaml:"port"`
// 	User   string `yaml:"user"`
// 	Pass   string `yaml:"pass"`
// 	DbName string `yaml:"db_name"`
// }

type storageConfig struct {
	DatabaseURL string `env:"DATABASE_URL"`
}

func Config() config {
	return cfg
}

func init() {
	envType := getEnvType()
	path := getConfigFilePath(envType)
	cleanenv.ReadConfig(path, &cfg)
}

func getConfigFilePath(envType string) string {
	path := fmt.Sprintf("./config/%s.yaml", envType)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("%s file not found", path)
	}
	return path
}

func getEnvType() string {
	envType := os.Getenv("ENV_TYPE")
	if envType == "" {
		log.Fatal("Empty ENV_TYPE variable")
		//envType = EnvLocal
	}
	if envType != EnvProd {
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
	}
	return envType
}
