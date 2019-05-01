/*
   Copyright 2017 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/codegangsta/cli.v1"
	"gopkg.in/jinzhu/gorm.v1"

	"github.com/gtalent/tendb/churchdirectory"
	"github.com/gtalent/tendb/importers"
	"github.com/gtalent/tendb/users"
)

func home() string {
	h := os.Getenv("GCDB_HOME")
	if h == "" {
		h = "."
	}
	h += "/"
	return h
}

func openDatabase() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=localhost user=postgres dbname=tendb sslmode=disable password=postgres")
}

func serve(c *cli.Context) error {
	_, err := openDatabase()
	if err != nil {
		fmt.Println(err)
		return err
	}
	s := &http.Server{
		Addr: "0.0.0.0:3000",
	}
	churchdirectory.SetupViews("/api/church_directory/", s)
	return s.ListenAndServe()
}

func migrate(c *cli.Context) error {
	db, err := openDatabase()
	if err != nil {
		fmt.Println(err)
		return err
	}
	users.Migrate(db)
	churchdirectory.Migrate(db)
	return nil
}

func importSK(c *cli.Context) error {
	db, err := openDatabase()
	if err != nil {
		fmt.Println(err)
		return err
	}
	path := c.String("path")
	err = importers.ImportSK(db, path)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func main() {
	churchdirectory.LsAssets()
	app := cli.NewApp()
	app.Name = "tendb"
	app.Usage = "Gary's Church Database"

	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "Run web server for tendb",
			Flags:  []cli.Flag{},
			Action: serve,
		},
		{
			Name:   "migrate",
			Usage:  "Migrate the database to the currently supported schema",
			Flags:  []cli.Flag{},
			Action: migrate,
		},
		{
			Name:  "createuser",
			Usage: "Creates a user",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "s,superuser",
					Usage: "Indicates whether or not the new user should be a superuser/admin",
				},
			},
			Action: func(c *cli.Context) {},
		},
		{
			Name:  "import-sk",
			Usage: "Import Servant Keeper 6 CSV file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p,path",
					Usage: "Path of Servant Keeper export directory",
				},
			},
			Action: importSK,
		},
	}

	app.Run(os.Args)
}
