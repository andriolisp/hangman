package infra

import (
	"github.com/andriolisp/hangman/infra/config"
	"github.com/andriolisp/hangman/infra/logger"
)

//BaseConfig has the signature which can be common used in the whole application
type BaseConfig interface {
	Config() *config.Config
	Log() *logger.Logger
}

// Base has all infra integrations
type Base struct {
	cfg *config.Config
	log *logger.Logger
}

// Config has all the application settings
func (b Base) Config() *config.Config {
	return b.cfg
}

// Log has all the log methods
func (b Base) Log() *logger.Logger {
	return b.log
}

// New return a new Infra Integration instance
func New() (*Base, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, err
	}

	log, err := logger.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Base{
		cfg: cfg,
		log: log,
	}, nil
}
