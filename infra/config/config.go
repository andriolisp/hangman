package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config has all the application settings
type Config struct {
	App AppConfig
	Log LogConfig
	DB  DBConfig
}

// DBConfig represents the configuration values about the DB.
type DBConfig struct {
	Name             string
	MaxLifeInMinutes int
}

// AppConfig has the base for application
type AppConfig struct {
	Name   string
	Debug  bool
	Port   int64
	Prefix string
}

// LogConfig represents the configuration values about the logging config.
type LogConfig struct {
	LogToFile bool
	Path      string
}

// setup set the initial configuration
func setup() {
	viper.SetEnvPrefix("api")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", ","))
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
}

// Read return the Config loaded config structure
func Read() (*Config, error) {
	setup()
	err := viper.ReadInConfig()
	if err != nil {
		if err == os.ErrNotExist {
			viper.SetConfigName("config.local")
			err = viper.ReadInConfig()
		}

		if err != nil {
			return nil, err
		}
	}

	return &Config{
		App: AppConfig{
			Name:   viper.GetString("app.name"),
			Debug:  viper.GetBool("app.debug"),
			Port:   viper.GetInt64("app.port"),
			Prefix: viper.GetString("app.prefix"),
		},
		Log: LogConfig{
			LogToFile: viper.GetBool("log.log-to-file"),
			Path:      viper.GetString("log.path"),
		},
		DB: DBConfig{
			Name:             viper.GetString("db.name"),
			MaxLifeInMinutes: viper.GetInt("db.max-life-minutes"),
		},
	}, nil
}
