package databolt

import (
	"sync"
	"time"

	"github.com/andriolisp/hangman/application/domain/contract"
	"github.com/andriolisp/hangman/infra"
	"github.com/andriolisp/hangman/infra/errors"
	"github.com/boltdb/bolt"
)

var (
	instance *Conn
	once     sync.Once
	connErr  error
)

// Conn has the structure to communicate with the database
type Conn struct {
	db   *bolt.DB
	base infra.BaseConfig
	game *gameRepo
}

// Close closes the db connection
func (c *Conn) Close() (err error) {
	return c.db.Close()
}

func getBaseConfig(base infra.BaseConfig) *bolt.Options {
	return &bolt.Options{
		Timeout: time.Duration(base.Config().DB.MaxLifeInMinutes) * time.Minute,
	}
}

// Instance return a new connection
func Instance(base infra.BaseConfig) (contract.DataManager, error) {
	once.Do(func() {
		db, err := bolt.Open(base.Config().DB.Name, 0600, getBaseConfig(base))
		if err != nil {
			connErr = errors.New(err.Error())
			return
		}

		instance = new(Conn)
		instance.db = db
		instance.base = base
		instance.game = newGameRepo(base, db)
	})

	return instance, connErr
}

//Game has all the methods to communicate with the Game bucket on the database
func (c *Conn) Game() contract.GameRepo {
	return c.game
}
