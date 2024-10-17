package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  bool
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	DB                 string
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "config/docker"
	} else if env == "production" {
		return "config/production"
	} else {
		return "config/development"
	}
}

func LoadingConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("unable to parse config, %v", err)
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to parse config, %v", err)
		return nil, err
	}
	return &config, nil
}

func GetConfig() *Config {
	path := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadingConfig(path, "yaml")
	if err != nil {
		log.Fatalf("unable to load config file, %v", err)
	}
	config, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("unable to parse config file, %v", err)
	}

	return config
}
