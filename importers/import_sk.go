/*
   Copyright 2017 gtalent2@gmail.com

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

	"github.com/gtalent/tendb/churchdirectory"
)

var genderKey = map[string]int{
	"Female": 0,
	"Male":   0,
}

type family struct {
	father   *churchdirectory.Person
	mother   *churchdirectory.Person
	children []*churchdirectory.Person
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
func ImportSK(db *gorm.DB, path string) error {
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
	var eventBaptism churchdirectory.EventType
	var eventDeath churchdirectory.EventType
	var eventLeftChurch churchdirectory.EventType
	var eventWedding churchdirectory.EventType
	var clearanceBackgroundCheck churchdirectory.ClearanceType
	db.Where("name = ?", churchdirectory.EventBaptism).First(&eventBaptism)
	db.Where("name = ?", churchdirectory.EventDeath).First(&eventDeath)
	db.Where("name = ?", churchdirectory.EventLeftChurch).First(&eventLeftChurch)
	db.Where("name = ?", churchdirectory.EventWedding).First(&eventWedding)
	db.Where("name = ?", churchdirectory.ClearanceBackgroundCheck).First(&clearanceBackgroundCheck)

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
		p := new(churchdirectory.Person)
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
			case churchdirectory.Male:
				fam.father = p
			case churchdirectory.Female:
				fam.mother = p
			}
		case "Daughter", "Son":
			fam.children = append(fam.children, p)
		}
		db.Create(&p)

		// setup events
		if d := parseDate(rec[BaptizedDate]); d != nil {
			e := churchdirectory.Event{EventType: eventBaptism, Person: *p, Date: *d}
			db.Create(&e)
		}
		if d := parseDate(rec[DeceasedDate]); d != nil {
			e := churchdirectory.Event{EventType: eventDeath, Person: *p, Date: *d}
			db.Create(&e)
		}
		if d := parseDate(rec[DateRemoved]); d != nil {
			e := churchdirectory.Event{EventType: eventLeftChurch, Person: *p, Date: *d}
			db.Create(&e)
		}
		if d := parseDate(rec[WeddingDate]); d != nil {
			e := churchdirectory.Event{EventType: eventWedding, Person: *p, Date: *d}
			db.Create(&e)
		}

		// set up background check
		if d := parseDate(rec[BackgroundCheck]); d != nil {
			e := churchdirectory.Clearance{ClearanceType: clearanceBackgroundCheck, Person: *p, Date: *d}
			db.Create(&e)
		}

		// setup roles
		roles := strings.Split(rec[Leadership], "; ")
		for _, v := range roles {
			var r = churchdirectory.Role{Name: v}
			db.FirstOrCreate(&r, r)
			ra := churchdirectory.RoleAssignment{Person: *p, Role: r}
			db.Create(&ra)
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
				fc := churchdirectory.ParentChildRelationship{Parent: *f.father, Child: *c}
				db.Create(&fc)
			}
			if f.mother != nil {
				mc := churchdirectory.ParentChildRelationship{Parent: *f.mother, Child: *c}
				db.Create(&mc)
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
