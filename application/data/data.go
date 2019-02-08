package data

import (
	"github.com/andriolisp/hangman/application/data/databolt"
	"github.com/andriolisp/hangman/application/domain/contract"
	"github.com/andriolisp/hangman/infra"
)

//Connect return a connection from a DB
func Connect(base infra.BaseConfig) (contract.DataManager, error) {
	return databolt.Instance(base)
}
