/*
   Copyright 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

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