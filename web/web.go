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

func SetupViews(prefix string) {
	if prefix[len(prefix)-1] != '/' {
		prefix += "/"
	}
	http.HandleFunc(prefix+"persons/", persons)
	http.HandleFunc(prefix+"users/", users)
}
