package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB Database
}

type Database struct {
	Driver   string
	Host     string
	Port     int
	DB       string
	Username string
	Password string
}

func Init(filePath, fileType string) (*Config, error) {
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		DB: Database{
			Driver:   viper.GetString("database.driver"),
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			DB:       viper.GetString("database.db"),
		},
	}, nil
}
