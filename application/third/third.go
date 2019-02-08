package third

import (
	"sync"

	"github.com/andriolisp/hangman/application/third/randomlist"
	"github.com/andriolisp/hangman/infra"
)

var (
	instance *Third
	once     sync.Once
)

// Third holds the all the third-party API's
type Third struct {
	RandomList *randomlist.RandomListService
}

// New returns a new domain Service instance
func New(base infra.BaseConfig) *Third {
	once.Do(func() {
		instance = &Third{
			RandomList: randomlist.NewRandomListService(base),
		}
	})

	return instance
}
