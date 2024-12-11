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
}

// LoadConfig loads config from file
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")    // Config file name without extension
	viper.SetConfigType("json")      // Config file format
	viper.AddConfigPath("../config") // path to look for the config file in
	viper.AddConfigPath("./config")  // path to look for the config file in
	viper.AddConfigPath(".")         // path to look for the config file in
	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal into Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
