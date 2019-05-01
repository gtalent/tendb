/*
   Copyright 2017 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package churchdirectory

import (
	"time"

	"gopkg.in/jinzhu/gorm.v1"
)

const (
	Female = 0
	Male   = 1
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

type ClearanceType struct {
	gorm.Model
	Name     string `gorm:"size:50"`
	Duration time.Duration
	Builtin  bool
}

type Clearance struct {
	gorm.Model
	ClearanceType      ClearanceType `gorm:"ForeignKey:ClearanceTypeRefer"`
	ClearanceTypeRefer uint64
	Person             Person
	Date               time.Time
}

// Role model to store roles in the database.
type Role struct {
	gorm.Model
	Name string `gorm:"size:50"`
}

// RoleAssignment model to store roles assignments in the database.
type RoleAssignment struct {
	gorm.Model
	Role      Role `gorm:"ForeignKey:RoleRefer"`
	RoleRefer uint64
	StartDate *time.Time
	EndDate   *time.Time
	Person    Person
}

// EventType model to store types of life events in the database.
type EventType struct {
	gorm.Model
	Name    string `gorm:"size:50;unique"`
	Builtin bool
}

// Event model to store life events in the database.
type Event struct {
	gorm.Model
	EventType      EventType `gorm:"ForeignKey:EventTypeRefer"`
	EventTypeRefer uint64
	Date           time.Time
	Person         Person
}

type ParentChildRelationship struct {
	gorm.Model
	Parent      Person `gorm:"ForeignKey:ParentRefer"`
	Child       Person `gorm:"ForeignKey:ChildRefer"`
	ParentRefer uint64
	ChildRefer  uint64
}

// Person model to store a person in the database.
type Person struct {
	gorm.Model
	FirstName    string `gorm:"size:50"`
	MiddleName   string `gorm:"size:50"`
	LastName     string `gorm:"size:50"`
	Suffix       string `gorm:"size:5"`
	Married      bool
	Sex          int
	Birthday     *time.Time
	HomePhone    string `gorm:"size:10"`
	CellPhone    string `gorm:"size:10"`
	EmailAddress string `gorm:"size:75"`
	AddressLine1 string `gorm:"size:50"`
	AddressLine2 string `gorm:"size:50"`
	City         string `gorm:"size:50"`
	Province     string `gorm:"size:50"`
	ZipCode      string `gorm:"size:10"`
	Homebound    bool
	OutOfArea    bool
	Member       bool
	Notes        string
	PicturePath  string `gorm:"size:255"`
}
