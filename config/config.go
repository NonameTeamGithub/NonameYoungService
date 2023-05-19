package config

import (
	"InternService/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	RedisDB struct {
	}
	MongoDB struct {
		DatabaseConnection string
		DatabaseName       string
	}
	PostgreSQLDB struct {
		User     string
		Pass     string
		Host     string
		Port     string
		Dbname   string
		SSLMode  string
		MaxConns string
	}
	TokenAccessExpiration int
}

func LoadConfig() *viper.Viper {
	log := logger.GetLogger()
	v := viper.New()
	v.AddConfigPath("../config/")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to load config.")
	}
	log.Info().Msg("Config loaded successfully.")
	return v
}

func ParseConfig(v *viper.Viper) Config {
	log := logger.GetLogger()
	var conf Config
	err := v.Unmarshal(&conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse config.")
	}
	log.Info().Msg("Config parsed successfully.")
	return conf
}
