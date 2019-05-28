/*
   Copyright 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package main

import (
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/gtalent/tendb/db"
	"github.com/gtalent/tendb/importers"
	"github.com/gtalent/tendb/web"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

func serve(c *cli.Context) error {
	const addr = "0.0.0.0:2010"
	web.SetupViews("/api/")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func migrate(c *cli.Context) error {
	conn := db.OpenDatabase()
	defer conn.Close()
	err := db.Migrate(conn)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func createUser(c *cli.Context) error {
	conn := db.OpenDatabase()
	defer conn.Close()
	var u db.User

	u.EmailAddress = c.String("email")
	u.FirstName = c.String("fn")
	u.LastName = c.String("ln")

	// get password from stdin
	print("Password: ")
	pw, err := terminal.ReadPassword(int(syscall.Stdin))
	println()
	if err != nil {
		return err
	}

	err = u.SetPassword(pw)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = conn.Insert(&u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	println("User " + u.EmailAddress + " created")
	return nil
}

func importSK(c *cli.Context) error {
	conn := db.OpenDatabase()
	defer conn.Close()
	path := c.String("path")
	err := importers.ImportSK(conn, path)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func main() {
	app := cli.NewApp()
	app.Name = "10db"
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
				cli.StringFlag{
					Name:  "e,email",
					Usage: "Email of account, which will also serve as the username",
				},
				cli.StringFlag{
					Name:  "fn",
					Usage: "First Name",
				},
				cli.StringFlag{
					Name:  "ln",
					Usage: "Last Name",
				},
			},
			Action: createUser,
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
