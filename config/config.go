package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB     Database
	Server Server
}

type Database struct {
	Driver   string
	Host     string
	Port     string
	DB       string
	Username string
	Password string
}

type Server struct {
	Host string
	Port string
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
			Port:     viper.GetString("database.port"),
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			DB:       viper.GetString("database.db"),
		},
		Server: Server{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
	}, nil
}
