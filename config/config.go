package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"time"
)

type Config struct {
	Postgres Postgres
	Server   Server
	Logger   Logger
}

type Server struct {
	Port int
}

type Logger struct {
	Level string
}

type Postgres struct {
	Host               string
	Port               string
	Username           string
	Password           string
	Database           string
	Debug              bool
	SslMode            string        `mapstructure:"ssl-mode"`
	MaxIdleConnections int           `mapstructure:"max-idle-connections"`
	MaxOpenConnections int           `mapstructure:"max-open-connections"`
	MaxConnectionTTL   time.Duration `mapstructure:"max-connection-ttl"`
	MigrationPath      string        `mapstructure:"migration-path"`
}

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app-config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config file")
	}
	return &conf, nil
}

func SetUpLogger(conf *Config) *logrus.Entry {
	logLevel, err := logrus.ParseLevel(conf.Logger.Level)
	if err != nil {
		logrus.WithError(err).Panic("Invalid logging level in configuration")
	}
	logrus.SetReportCaller(true)
	logrus.SetLevel(logLevel)
	return logrus.NewEntry(logrus.StandardLogger())
}
