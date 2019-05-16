package cmd

import (
	"fmt"
	"mess-app/internal/core"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type configs struct {
	dbConfig struct {
		host     string
		user     string
		password string
		dbName   string
		port     int
	}
	logger struct {
		level  logrus.Level
		output *os.File
	}
	appConfig struct {
		port string
		host string
	}
}

func mustReadConfig() (*configs, error) {

	viper.SetConfigName("app")
	viper.SetConfigType("json")
	viper.AddConfigPath(core.CONFIG_FILE)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "Failed to parse config file")
	}
	c := configs{}
	c.appConfig.host = viper.GetString("app.appconfig.host")
	c.appConfig.port = viper.GetString("app.appconfig.port")

	c.dbConfig.host = viper.GetString("app.db.host")
	c.dbConfig.user = viper.GetString("app.db.user")
	c.dbConfig.password = viper.GetString("app.db.password")
	c.dbConfig.dbName = viper.GetString("app.db.dbname")
	c.dbConfig.port = viper.GetInt("app.db.port")

	switch viper.GetString("app.logger.level") {
	case core.LOG_DEBUG:
		c.logger.level = logrus.DebugLevel
	case core.LOG_INFO:
		c.logger.level = logrus.InfoLevel
	case core.LOG_WARNING:
		c.logger.level = logrus.WarnLevel
	case core.LOG_ERROR:
		c.logger.level = logrus.ErrorLevel
	}
	switch viper.GetString("app.logger.output") {
	case core.OUT_STDOUT:
		c.logger.output = os.Stdout
	}
	return &c, nil
}

func mustValidateConfig(c *configs) error {
	if c.dbConfig.dbName == "" || c.dbConfig.host == "" || c.dbConfig.port == 0 || c.dbConfig.user == "" {
		return errors.New(fmt.Sprintf("Invalid db configuration, DB config : %+v", c.dbConfig))
	}
	if c.appConfig.host == "" || c.appConfig.port == "" {
		return errors.New(fmt.Sprintf("Invalid appconfig, App config %+v", c.appConfig))
	}
	if c.logger.level == 0 || c.logger.output == nil {
		return errors.New(fmt.Sprintf("Invalid logger configuration, log config %+v", c.logger))
	}
	return nil
}
