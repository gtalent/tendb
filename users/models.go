/*
   Copyright 2017 - 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package users

import (
	"golang.org/x/crypto/bcrypt"

	"gopkg.in/jinzhu/gorm.v1"
)

// User model to store a user in the database.
type User struct {
	gorm.Model
	EmailAddress string `gorm:"size:75"`
	PasswordHash []byte `gorm:"size:75"`
}

/*
SetPassword hashes and sets the user password.
Returns error if there was an error.
*/
func (u *User) SetPassword(pw string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = hash
	return nil
}

// MigrateUser migrates changes on User table based on User type.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
