package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type config struct {
	Path string
}

var conf *config

func SetConfigPath(path string) {
	conf = &config{Path: path}
}

func Read() *Config {
	var (
		cfg *Config
		err error
	)

	viper.SetConfigType("env")
	viper.SetConfigFile(conf.Path)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		log.Fatal().Msgf("failed to read config. error: %s", err.Error())
	}

	cfg = &Config{}
	err = viper.Unmarshal(cfg)

	if err != nil {
		log.Fatal().Msgf("failed to unmarshall config. error: %s", err.Error())
	}

	return cfg
}
