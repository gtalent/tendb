/*
   Copyright 2017 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package churchdirectory

import (
	"gopkg.in/jinzhu/gorm.v1"
)

// Migrate performs the database migrations needed by this package.
func Migrate(db *gorm.DB) {
	// setup tables
	db.AutoMigrate(&EventType{})
	db.AutoMigrate(&ClearanceType{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Clearance{})
	db.AutoMigrate(&Event{})
	db.AutoMigrate(&RoleAssignment{})
	db.AutoMigrate(&ParentChildRelationship{})

	// enter default values

	// default EventTypes
	for _, v := range DefaultEventTypes {
		e := EventType{
			Name:    v,
			Builtin: true,
		}
		db.FirstOrCreate(&e, e)
	}

	// default ClearanceTypes
	for _, v := range DefaultClearanceTypes {
		e := ClearanceType{
			Name:    v,
			Builtin: true,
		}
		db.FirstOrCreate(&e, e)
	}
}
