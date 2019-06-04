/*
   Copyright 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package main

import "github.com/go-pg/pg"

type config struct {
	DbOptions pg.Options `json:"db_options"`
}
