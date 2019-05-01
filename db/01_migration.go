
package db

import (
	"github.com/go-pg/migrations"
)

func init() {
	const dir = "01_init"
	migrations.MustRegisterTx(func(db migrations.DB) error {
		path := dir + "/up.sql"
		sql, err := loadMigrationFile(path)
		if err != nil {
			return err
		}
		_, err = db.Exec(sql)
		return err
	}, func(db migrations.DB) error {
		path := dir + "/down.sql"
		sql, err := loadMigrationFile(path)
		if err != nil {
			return err
		}
		_, err = db.Exec(sql)
		return err
	})
}