package config

/* *****Test Project for HeliTechnology Company*****
...... xhmaozedi@gmail.com .......
*/
import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	DB struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"db"`
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	S3Config struct {
		Endpoint  string `mapstructure:"endpoint"`
		Bucket    string `mapstructure:"bucket"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"s3_config"`
	MinioConfig struct {
		Host      string `mapstructure:"end_point"`
		Port      string `mapstructure:"port"`
		Bucket    string `mapstructure:"bucket"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
		TLS       bool   `mapstructure:"tls"`
	} `mapstructure:"minio_config"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
