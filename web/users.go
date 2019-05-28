/*
   Copyright 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/gtalent/tendb/db"
	"log"
)

func users(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		usersCreate(rw, r)
	case "GET":
		usersRead(rw, r)
	case "PUT":
		usersUpdate(rw, r)
	case "DELETE":
		usersDelete(rw, r)
	}
}

func usersCreate(rw http.ResponseWriter, r *http.Request) {
	type createUserMsg struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	var cu createUserMsg
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(b, &cu)
	if err != nil {
		log.Println(err)
		return
	}

	// create user in database
	var u db.User
	u.EmailAddress = cu.Email
	u.FirstName = cu.FirstName
	u.LastName = cu.LastName
	u.SetPassword([]byte(cu.Password))
	conn := db.OpenDatabase()
	err = conn.Insert(&u)
	if err != nil {
		log.Println(err)
		return
	}
}

func usersRead(rw http.ResponseWriter, r *http.Request) {
}

func usersUpdate(rw http.ResponseWriter, r *http.Request) {
}

func usersDelete(rw http.ResponseWriter, r *http.Request) {
}
