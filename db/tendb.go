/*
   Copyright 2017 - 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package db

import "github.com/jinzhu/gorm"

func OpenDatabase() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=localhost user=postgres dbname=tendb sslmode=disable password=postgres")
}
