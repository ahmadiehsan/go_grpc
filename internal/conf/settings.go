package conf

import (
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Settings struct {
	Debug    bool   `mapstructure:"DEBUG"`
	DBName   string `mapstructure:"DB_NAME"`
	HTTPPort string `mapstructure:"HTTP_PORT"`
}

func LoadSettings() (s Settings) {
	configPath, err := filepath.Abs("configs")
	if err != nil {
		log.Fatal().Err(err).Msg("Can't find the absolute path of configs directory")
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env") // TODO in docker-compose we don't need this
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Can't read settings")
	}

	err = viper.Unmarshal(&s)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't unmarshal settings")
	}

	return
}
