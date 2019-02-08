package databolt

import (
	"fmt"

	"github.com/andriolisp/hangman/application/entity"
	"github.com/andriolisp/hangman/infra"
	"github.com/boltdb/bolt"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/satori/go.uuid"
)

type gameRepo struct {
	conn executor
	infra.BaseConfig
	bucket *bolt.Bucket
}

func newGameRepo(base infra.BaseConfig, conn executor) *gameRepo {
	var bucket *bolt.Bucket
	var err error

	conn.Update(func(tx *bolt.Tx) error {
		bucket, err = tx.CreateBucketIfNotExists([]byte("game"))
		if err != nil {
			base.Log().Error("Error to create game bucket: ", err)
			tx.Rollback()
			return err
		} else {
			return nil
		}
	})

	return &gameRepo{
		conn,
		base,
		bucket,
	}
}

//Get will return an specific game
func (r *gameRepo) Get(ID string) (*entity.Game, error) {
	if ID != "" {
		game := new(entity.Game)
		var bGame []byte

		r.conn.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("game"))
			bGame = b.Get([]byte(ID))
			return nil
		})

		if err := ffjson.Unmarshal(bGame, game); err != nil {
			fmt.Println(err)
			return nil, err
		}

		return game, nil
	}
	return nil, nil
}

//Put will update an existing game
func (r *gameRepo) Put(game *entity.Game) (*entity.Game, error) {
	if game.ID == "" {
		return nil, nil
	} else {
		bGame, err := ffjson.Marshal(game)
		if err != nil {
			return nil, err
		}

		if err := r.conn.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("game"))
			return bucket.Put([]byte(game.ID), bGame)
		}); err != nil {
			return nil, err
		}

		return r.Get(game.ID)
	}
}

//Add will create a new game structure
func (r *gameRepo) Add(game *entity.Game) (*entity.Game, error) {
	if game.ID != "" {
		return r.Get(game.ID)
	} else {
		id, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		game.ID = id.String()

		return r.Put(game)
	}
}
