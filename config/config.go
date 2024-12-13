package config

/* *****Test Project for HeliTechnology Company*****
...... xhmaozedi@gmail.com .......
*/
import (
	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	Server struct {
		Port int
	}
	S3Config struct {
		Endpoint  string
		Bucket    string
		AccessKey string
		SecretKey string
	}
	MinioConfig struct {
		Host      string
		Port      string
		Bucket    string
		AccessKey string
		SecretKey string
		TLS       bool
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
