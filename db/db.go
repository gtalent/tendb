
package db

import (
	"github.com/go-pg/pg"
)

func OpenDatabase() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "tendb",
	})
}
