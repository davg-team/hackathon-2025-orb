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
	EnvType string        `yaml:"env_type"`
	Server  serverConfig  `yaml:"server"`
	Storage storageConfig `yaml:"storage"`
	Client  clientConfig  `yaml:"client"`
	Bucket  bucketConfig  `yaml:"bucket"`
}

type serverConfig struct {
	Port int `yaml:"port"`
}

// type bucketConfig struct {
// 	Endpoint   string `yaml:"endpoint"`
// 	AccessKey  string `yaml:"access_key"`
// 	SecretKey  string `yaml:"secret_key"`
// 	BucketName string `yaml:"bucket_name"`
// }

type bucketConfig struct {
	Endpoint   string `env:"MINIO_ENDPOINT"`
	AccessKey  string `env:"MINIO_ACCESS_KEY"`
	SecretKey  string `env:"MINIO_SECRET_KEY"`
	BucketName string `env:"MINIO_BUCKET_NAME"`
}

// type clientConfig struct {
// 	LayerID  string `yaml:"layer_id"`
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// }

type clientConfig struct {
	LayerID  string `env:"CLIENT_LAYER_ID"`
	Username string `env:"CLIENT_USERNAME"`
	Password string `env:"CLIENT_PASSWORD"`
}

type storageConfig struct {
	URL string `env:"DATABASE_URL"`
}

// Cfg return copy of cfg (line 18)
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
		// log.Fatal("Empty ENV_TYPE variable")
		envType = EnvLocal
	}
	if envType != EnvProd {
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
	}
	return envType
}
