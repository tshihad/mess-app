package cmd

import (
	"mess-app/internal/core"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("json")
	viper.AddConfigPath(core.CONFIG_FILE)
	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Wrap(err, "Failed to parse config file"))
	}
}
