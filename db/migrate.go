/*
   Copyright 2017 - 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package db

//go:generate go run ../assets_generate/assets_generate.go db

import (
	"io"

	"github.com/go-pg/pg"
	"github.com/go-pg/migrations"
)

func loadMigrationFile(path string) (string, error) {
	file, err := assets.Open(path)
	if err != nil {
		return "", err
	}
	stat, _ := file.Stat()
	buff := make([]byte, stat.Size())
	_, err = io.ReadFull(file, buff)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

// Migrate performs the database migrations needed by this package.
func Migrate(db *pg.DB) error {
	_, _, err := migrations.Run(db, "init")
	_, _, err = migrations.Run(db, "up")
	return err
}
