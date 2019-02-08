package applicationservice

import (
	"sync"

	"github.com/andriolisp/hangman/application/service"
	"github.com/andriolisp/hangman/application/third"
	"github.com/andriolisp/hangman/infra"
)

var (
	instance *Service
	once     sync.Once
	appError error
)

//Service will return all the methods to get the user information
type Service struct {
	svc *service.Service
	infra.BaseConfig
	third *third.Third
	Game  *gameApp
}

//New return an instance of the Service struct
func New(base infra.BaseConfig) (*Service, error) {
	once.Do(func() {
		svc, err := service.New(base)
		if err != nil {
			appError = err
			return
		}

		instance = &Service{svc, base, nil, nil}
		instance.third = third.New(base)
		instance.Game = newGameApp(instance, base, instance.third)
	})

	return instance, appError
}
