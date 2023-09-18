package conf

import (
	"github.com/nehaal10/ecommerce-api/internal/utils"
	"github.com/spf13/viper"
)

type Config struct {
	DbHost string `mapstructure:"MONGO_URL"`
	Port   string `mapstructure:"PORT"`
	Secret string `mapstructure:"SECRET_KEY"`
}

func NewConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	utils.Checkerr(err)

	viper.Unmarshal(&config)
	return
}
