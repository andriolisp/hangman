package gamerouter

import (
	"strings"
	"sync"

	"github.com/andriolisp/hangman/application/entity"
	"github.com/andriolisp/hangman/infra"
)

var (
	healerOnce   sync.Once
	helperStruct *helper
)

type helper struct {
	infra.BaseConfig
}

func helperInstance(base infra.BaseConfig) *helper {
	healerOnce.Do(func() {
		helperStruct = &helper{
			base,
		}
	})

	return helperStruct
}

func (h *helper) GameHelper(game *entity.Game) *entity.Game {
	if game != nil {
		for _, s := range strings.Split(game.Word, "") {
			if len(game.Details) > 0 && game.Details.HasLetter(s) || game.Winner != 0 {
				game.Replacers = append(game.Replacers, s)
			} else {
				game.Replacers = append(game.Replacers, "_")
			}
		}

		game.Word = ""
	}
	return game
}
