package badger

import (
	badger "github.com/dgraph-io/badger/v4"
)

type Config struct {
	Dir string
}

func New(c *Config) *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(c.Dir))
	if err != nil {
		panic(err)
	}

	return db
}
