/*
   Copyright 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package db

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	EventBaptism      = "Baptism"
	EventDeath        = "Death"
	EventJoinedChurch = "Joined Church"
	EventLeftChurch   = "Left Church"
	EventWedding      = "Wedding"

	ClearanceBackgroundCheck = "Background Check"
)

var DefaultEventTypes = []string{
	EventBaptism,
	EventDeath,
	EventJoinedChurch,
	EventLeftChurch,
	EventWedding,
}

var DefaultClearanceTypes = []string{
	ClearanceBackgroundCheck,
}

/*
SetPassword hashes and sets the user password.
Returns error if there was an error.
*/
func (u *User) SetPassword(pw []byte) error {
	hash, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = &hash
	return nil
}
