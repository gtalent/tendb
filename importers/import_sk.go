/*
   Copyright 2017 - 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package importers

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/jinzhu/gorm.v1"

	"github.com/gtalent/tendb/db"
)

var genderKey = map[string]int{
	"Female": 0,
	"Male":   0,
}

type family struct {
	father   *db.Person
	mother   *db.Person
	children []*db.Person
}

func buildKey(h []string) map[string]int {
	k := make(map[string]int)
	for i, v := range h {
		k[v] = i
	}
	return k
}

func member(ms string) bool {
	switch ms {
	case "Active Member", "Homebound Member", "Out-of-the-area Member":
		return true
	default:
		return false
	}
}

func parseDate(d string) *time.Time {
	t, _ := time.Parse(time.RFC3339, d+"T00:00:00Z")
	return &t
}

func cleanPhoneNumber(n string) string {
	out := ""
	for _, v := range n {
		if v <= '0' && v >= '9' {
			out += string(v)
		}
	}
	return out
}

// ImportSK imports a Servant Keeper data directory.
func ImportSK(conn *gorm.DB, path string) error {
	// load CSV data
	f, err := os.Open(path + "/sk.csv")
	if err != nil {
		return err
	}
	reader := csv.NewReader(f)

	rec, err := reader.Read()
	if err != nil {
		return err
	}

	key := buildKey(rec)
	FirstName := key["First Name"]
	MiddleName := key["Middle Name"]
	LastName := key["Last Name"]
	Suffix := key["Suffix"]
	MaritalStatus := key["Marital Status"]
	HomePhone := key["HomePhone"]
	CellPhone := key["CellPhone"]
	Email := key["E-Mail"]
	AddressLine1 := key["Address"]
	AddressLine2 := key["Address Line 2"]
	City := key["City"]
	ZipCode := key["Zip Code"]
	State := key["State"]
	MemberStatus := key["Member Status"]
	BirthDate := key["Birth Date"]
	Relationship := key["Relationship"]
	DirectoryName := key["Directory Name"]
	Gender := key["Gender"]
	BaptizedDate := key["Baptized Date"]
	DeceasedDate := key["Deceased date"]
	DateRemoved := key["Date Removed"]
	WeddingDate := key["Wedding Date"]
	BackgroundCheck := key["Bckg Check"]
	Leadership := key["Leadership"]
	var eventBaptism db.EventType
	var eventDeath db.EventType
	var eventLeftChurch db.EventType
	var eventWedding db.EventType
	var clearanceBackgroundCheck db.ClearanceType
	conn.Where("name = ?", db.EventBaptism).First(&eventBaptism)
	conn.Where("name = ?", db.EventDeath).First(&eventDeath)
	conn.Where("name = ?", db.EventLeftChurch).First(&eventLeftChurch)
	conn.Where("name = ?", db.EventWedding).First(&eventWedding)
	conn.Where("name = ?", db.ClearanceBackgroundCheck).First(&clearanceBackgroundCheck)

	families := make(map[string]*family)
	i := 0
	for rec, err = reader.Read(); err != io.EOF; rec, err = reader.Read() {
		if err != nil {
			continue
		}
		if i%50 == 0 {
			print("\r" + strconv.Itoa(i))
		}
		i++
		p := new(db.Person)
		p.FirstName = rec[FirstName]
		p.MiddleName = rec[MiddleName]
		p.LastName = rec[LastName]
		p.Suffix = rec[Suffix]
		p.Married = rec[MaritalStatus] == "Married"
		p.HomePhone = cleanPhoneNumber(rec[HomePhone])
		p.CellPhone = cleanPhoneNumber(rec[CellPhone])
		p.EmailAddress = rec[Email]
		p.AddressLine1 = rec[AddressLine1]
		p.AddressLine2 = rec[AddressLine2]
		p.ZipCode = rec[ZipCode]
		p.City = rec[City]
		p.Province = rec[State]
		p.Member = member(rec[MemberStatus])
		p.Birthday = parseDate(rec[BirthDate])
		p.Sex = genderKey[rec[Gender]]
		dirName := rec[DirectoryName]
		rel := rec[Relationship]
		fam, ok := families[dirName]
		if !ok {
			fam = new(family)
			families[dirName] = fam
		}
		switch rel {
		case "head of household", "Spouse", "Father", "Mother":
			switch p.Sex {
			case db.Male:
				fam.father = p
			case db.Female:
				fam.mother = p
			}
		case "Daughter", "Son":
			fam.children = append(fam.children, p)
		}
		conn.Create(&p)

		// setup events
		if d := parseDate(rec[BaptizedDate]); d != nil {
			e := db.Event{EventType: eventBaptism, Person: *p, Date: *d}
			conn.Create(&e)
		}
		if d := parseDate(rec[DeceasedDate]); d != nil {
			e := db.Event{EventType: eventDeath, Person: *p, Date: *d}
			conn.Create(&e)
		}
		if d := parseDate(rec[DateRemoved]); d != nil {
			e := db.Event{EventType: eventLeftChurch, Person: *p, Date: *d}
			conn.Create(&e)
		}
		if d := parseDate(rec[WeddingDate]); d != nil {
			e := db.Event{EventType: eventWedding, Person: *p, Date: *d}
			conn.Create(&e)
		}

		// set up background check
		if d := parseDate(rec[BackgroundCheck]); d != nil {
			e := db.Clearance{ClearanceType: clearanceBackgroundCheck, Person: *p, Date: *d}
			conn.Create(&e)
		}

		// setup roles
		roles := strings.Split(rec[Leadership], "; ")
		for _, v := range roles {
			var r = db.Role{Name: v}
			conn.FirstOrCreate(&r, r)
			ra := db.RoleAssignment{Person: *p, Role: r}
			conn.Create(&ra)
		}
	}
	println("\r" + strconv.Itoa(i))

	// setup families
	i = 0
	statusPrefix := "\rParent-child relationship: "
	for k, f := range families {
		f = families[k]
		for _, c := range f.children {
			if f.father != nil {
				fc := db.ParentChildRelationship{Parent: *f.father, Child: *c}
				conn.Create(&fc)
			}
			if f.mother != nil {
				mc := db.ParentChildRelationship{Parent: *f.mother, Child: *c}
				conn.Create(&mc)
			}
			if i%50 == 0 {
				print(statusPrefix + strconv.Itoa(i))
			}
			i++
		}
	}
	println(statusPrefix + strconv.Itoa(i))

	return nil
}
