package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Database     DatabaseConfig
		Server       ServerConfig
		AWS          AWSConfig
		Notification NotificationConfig
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}

	ServerConfig struct {
		Port  string
		Debug bool
	}

	AWSConfig struct {
		Region string
		S3     S3Config
	}

	S3Config struct {
		Bucket string
	}

	NotificationConfig struct {
		ApiKey     string
		TemplateID string
		From       string
	}
)

func LoadConfig(env string) (Config, error) {
	viper.SetConfigFile(fmt.Sprintf("config/config.%s.yaml", env))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error Unmarshal Config: %v\n", err)
		return Config{}, err
	}

	return config, nil
}
