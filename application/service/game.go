package service

import (
	"github.com/andriolisp/hangman/application/domain/contract"
	"github.com/andriolisp/hangman/application/entity"
	"github.com/andriolisp/hangman/infra"
)

type gameService struct {
	db     contract.DataManager
	parent *Service
	base   infra.BaseConfig
}

func newGameService(base infra.BaseConfig, svc *Service) *gameService {
	u := &gameService{
		svc.db,
		svc,
		base,
	}

	return u
}

func (s *gameService) Add(game *entity.Game) (*entity.Game, error) {
	return s.db.Game().Add(game)
}

func (s *gameService) Put(game *entity.Game) (*entity.Game, error) {
	return s.db.Game().Put(game)
}

func (s *gameService) Get(ID string) (*entity.Game, error) {
	return s.db.Game().Get(ID)
}
