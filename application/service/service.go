package service

import (
	"sync"

	"github.com/andriolisp/hangman/application/data"
	"github.com/andriolisp/hangman/application/domain/contract"
	"github.com/andriolisp/hangman/infra"
)

var (
	serviceErr error
	instance   *Service
	once       sync.Once
)

//Service has all the methods
type Service struct {
	db   contract.DataManager
	game *gameService
	base infra.BaseConfig
}

//New return a new instance of service
func New(base infra.BaseConfig) (*Service, error) {
	once.Do(func() {
		db, err := data.Connect(base)
		if err != nil {
			base.Log().Errorln("Error connection to the database: ", err)
			serviceErr = err
			return
		}
		base.Log().Printf("Connected to the database at %s.", base.Config().DB.Name)

		instance = &Service{
			db:   db,
			base: base,
		}
		instance.game = newGameService(base, instance)
	})

	return instance, serviceErr
}

//Game return an instance of the methods for game
func (s *Service) Game() *gameService {
	return s.game
}
