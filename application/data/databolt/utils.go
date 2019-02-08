package databolt

import "github.com/boltdb/bolt"

type executor interface {
	Update(fn func(*bolt.Tx) error) error
	View(fn func(*bolt.Tx) error) error
}
