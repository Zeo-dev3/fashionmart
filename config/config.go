package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
	Jwt      JwtConfig
	Swagger  SwaggerConfig
	Logger   LoggerConfig
}

type PostgresConfig struct {
	Port     int
	Host     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	AppName string
	Port    string
	Prefork bool
	Mode    string
}

type JwtConfig struct {
	SecretKey string
}

type SwaggerConfig struct {
	BasePath string
	FilePath string
	Path     string
	Title    string
}

type LoggerConfig struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

func LoadConfig(fileName string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(fileName)
	v.SetConfigType("yml")
	v.AddConfigPath("./")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
