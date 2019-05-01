/*
   Copyright 2017 - 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/codegangsta/cli.v1"

	"github.com/gtalent/tendb/churchdirectory"
	"github.com/gtalent/tendb/db"
	//"github.com/gtalent/tendb/importers"
	//"github.com/gtalent/tendb/users"
	"github.com/go-pg/pg"
)

func home() string {
	h := os.Getenv("GCDB_HOME")
	if h == "" {
		h = "."
	}
	h += "/"
	return h
}

func openDatabase() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "tendb",
	})
}

func serve(c *cli.Context) error {
	s := &http.Server{
		Addr: "0.0.0.0:3000",
	}
	churchdirectory.SetupViews("/api/church_directory/", s)
	return s.ListenAndServe()
}

func migrate(c *cli.Context) error {
	conn := openDatabase()
	defer conn.Close()
	err := db.Migrate(conn)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func importSK(c *cli.Context) error {
	//db := openDatabase()
	//defer conn.Close()
	//path := c.String("path")
	//err = importers.ImportSK(db, path)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//return err
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "tendb"
	app.Usage = "10db Church Database"

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
