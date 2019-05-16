package cmd

import (
	"mess-app/internal/core"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func readDBConfig() (map[string]string, error) {
	f, err := os.Open(core.CONFIG_FILE)
	defer f.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read config file")
	}
	if err := viper.ReadConfig(f); err != nil {
		return nil, errors.Wrap(err, "Failed to parse config file")
	}
	return nil, nil
}
