package config

import (
	"github.com/spf13/viper"
)

var Cfg Config

type (
	Config struct {
		App   App   `mapstructure:"app" validate:"required"`
		Redis Redis `mapstructure:"redis" validate:"required"`
	}

	App struct {
		Env         string `mapstructure:"env" validate:"required"`
		Port        string `mapstructure:"port" validate:"required"`
	}

	Redis struct {
		Address            string        `mapstructure:"REDIS_ADDRESS"`
	}
)

func setDefaultValues(v *viper.Viper) {
	v.SetDefault("app.env", "dev")
	v.SetDefault("app.port", "8081")
	v.SetDefault("redis.REDIS_ADDRESS", "localhost:6379")

}
