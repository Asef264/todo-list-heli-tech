package config

/* *****Test Project for HeliTechnology Company*****
...... xhmaozedi@gmail.com .......
*/
import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var AppConfig *config // global app config

type config struct {
	Database Postgres
	AWS      AwsS3Config
}

type Postgres struct {
	Host         string        `mapstructure:"host"`          // postgres host
	Port         string        `mapstructure:"port"`          // postgres port
	User         string        `mapstructure:"user"`          // postgres user
	Pass         string        `mapstructure:"pass"`          // postgres pass
	DatabaseName string        `mapstructure:"database_name"` // postgres database
	SslMode      string        `mapstructure:"ssl_mode"`      // postgres ssl mode
	Timeout      time.Duration `mapstructure:"timeout"`       // postgres timeout
}

type AwsS3Config struct {
	Endpoint  string `mapstructure:"endpoint"`
	Bucket    string `mapstructure:"bucket"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

// LoadConfig loads config from file
func LoadConfig(path string, isTest bool) {
	if isTest {
		viper.SetConfigName("config_test") // name of test config file (without extension)
	} else {
		viper.SetConfigName("config") // name of config file (without extension)
	}
	viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name

	if path == "" {
		viper.AddConfigPath("./config") // path to look for the config file in
	} else {
		viper.SetConfigFile(path)
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	AppConfig = &config{}
	if err = viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
