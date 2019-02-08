package contract

import "github.com/andriolisp/hangman/application/entity"

type repoManager interface {
	Game() GameRepo
}

//GameRepo has all signature methods for Game Interactions
type GameRepo interface {
	Get(ID string) (*entity.Game, error)
	Put(game *entity.Game) (*entity.Game, error)
	Add(game *entity.Game) (*entity.Game, error)
}
