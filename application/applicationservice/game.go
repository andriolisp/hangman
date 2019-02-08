package applicationservice

import (
	"fmt"
	"strings"

	"github.com/andriolisp/hangman/application/entity"
	"github.com/andriolisp/hangman/application/service"
	"github.com/andriolisp/hangman/application/third"
	"github.com/andriolisp/hangman/infra"
)

type gameApp struct {
	svc    *service.Service
	parent *Service
	third  *third.Third
	infra.BaseConfig
}

func newGameApp(app *Service, base infra.BaseConfig, third *third.Third) *gameApp {
	return &gameApp{
		app.svc,
		app,
		third,
		base,
	}
}

func (s *gameApp) Add(game *entity.Game) (*entity.Game, error) {
	game.Word = s.third.RandomList.GetRandomWord()
	game.WordSize = len(game.Word)
	game.Remaining = game.WordSize
	game.Turn = 1
	game.Winner = 0

	game.Players = make(map[string]*entity.Player)
	for i := 0; i < game.PlayersNum; i++ {
		turn := false
		if i == 0 {
			turn = true
		}

		player := entity.Player{
			Num:    i + 1,
			Points: 0,
			Dead:   false,
			Turn:   turn,
		}

		game.Players[fmt.Sprintf("player%v", i+1)] = &player
	}

	return s.svc.Game().Add(game)
}

func (s *gameApp) Put(game *entity.Game, detail *entity.Detail) (*entity.Game, error) {
	if game != nil && detail != nil {
		if game.Winner == -1 {
			game.Message = "All the players are already dead. Start a new game."
			return game, nil
		}

		letter := strings.TrimSpace(strings.ToUpper(detail.Letter))
		if len(letter) > 1 {
			game.Message = fmt.Sprintf("Player %v should add one letter at the time", game.Turn)
			return game, nil
		}

		if game.Details.HasLetter(letter) {
			game.Message = fmt.Sprintf("The letter %s has been chosen before", letter)
			return game, nil
		}

		detail.Player = game.Turn
		player := fmt.Sprintf("player%v", game.Turn)
		lettersFound := strings.Count(game.Word, letter)
		if lettersFound > 0 {
			detail.Points = lettersFound * 10
			detail.Found = true
			game.Players[player].Points += detail.Points
			game.Remaining = game.Remaining - lettersFound
			game.Message = fmt.Sprintf("Letter %s found by Player %v.", letter, detail.Player)
		} else {
			game.Players[player].Tentatives++
			if game.Players[player].Tentatives >= 5 {
				game.Players[player].Dead = true
			}
			game.Message = fmt.Sprintf("Letter %s not found.", letter)
		}

		if game.Remaining == 0 {
			game.Winner = game.Turn
			game.Message = fmt.Sprintf("Player %v won.", game.Turn)
		} else {
			if !s.allPlayersDead(game) {
				game.Players[player].Turn = false

				s.nextPlayer(game)
				for game.Players[fmt.Sprintf("player%v", game.Turn)].Dead {
					s.nextPlayer(game)
				}

				game.Players[fmt.Sprintf("player%v", game.Turn)].Turn = true
			} else {
				game.Message = "All players are dead."
				game.Winner = -1
			}
		}

		detail.Sequential = len(game.Details) + 1
		game.Details = append(game.Details, *detail)

		return s.svc.Game().Put(game)
	} else {
		return nil, nil
	}
}

func (s *gameApp) allPlayersDead(game *entity.Game) bool {
	for _, v := range game.Players {
		if !v.Dead {
			return false
		}
	}
	return true
}

func (s *gameApp) nextPlayer(game *entity.Game) {
	if game.Turn+1 > game.PlayersNum {
		game.Turn = 1
	} else {
		game.Turn = game.Turn + 1
	}
}

func (s *gameApp) Get(ID string) (*entity.Game, error) {
	return s.svc.Game().Get(ID)
}
