/*
   Copyright 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package web

import (
	"net/http"
)


func persons(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		personsGet(rw, r)
	}
}

func personsGet(rw http.ResponseWriter, r *http.Request) {
}
