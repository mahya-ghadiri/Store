package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func Init(path string) {
	v := viper.New()
	setDefaultValues(v)
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		logrus.Info("No config file found.", err)
	}

	if err := v.UnmarshalExact(&Cfg); err != nil {
		logrus.Panicf("invalid configuration: %s", err)
	}

}
