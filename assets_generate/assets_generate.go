/*
   Copyright 2017 - 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/shurcooL/vfsgen"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const tmpl = `
package PKG_NAME

import (
	"github.com/go-pg/migrations"
)

func init() {
	const dir = "DIR"
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
}`

func ls(dirPath string) ([]string, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	list, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, f := range list {
		println(f.Name())
		out = append(out, f.Name())
	}
	return out, nil
}

func main() {
	const path = "migrations"
	flag.Parse()
	pkg := flag.Arg(0)
	// bundle files
	err := vfsgen.Generate(http.Dir(path), vfsgen.Options{
		PackageName: pkg,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// generate migration files
	dirs, err := ls(path)
	if err != nil {
		log.Fatalln(err)
	}
	for i, dir := range dirs {
		out := strings.Replace(tmpl, "DIR", dir, 1)
		out = strings.Replace(out, "PKG_NAME", pkg, 1)
		fn := fmt.Sprintf("%02d", i+1) + "_migration.go"
		ioutil.WriteFile(fn, []byte(out), 0644)
	}
}
