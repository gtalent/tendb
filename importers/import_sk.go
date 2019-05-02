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

	"github.com/go-pg/pg"
	"github.com/gtalent/tendb/db"
)

var genderKey = map[string]int{
	"Unknown": db.Unknown,
	"Female":  db.Female,
	"Male":    db.Male,
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

func cleanPhoneNumber(n string) *string {
	out := ""
	for _, v := range n {
		if v <= '0' && v >= '9' {
			out += string(v)
		}
	}
	return &out
}

// ImportSK imports a Servant Keeper data directory.
func ImportSK(conn *pg.DB, path string) error {
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
	//conn.Model(&eventBaptism).Where("name = ?", db.EventBaptism).Select()
	//conn.Model(&eventDeath).Where("name = ?", db.EventDeath).Select()
	//conn.Model(&eventLeftChurch).Where("name = ?", db.EventLeftChurch).Select()
	//conn.Model(&eventWedding).Where("name = ?", db.EventWedding).Select()
	//conn.Model(&clearanceBackgroundCheck).Where("name = ?", db.ClearanceBackgroundCheck).Select()

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
		p.MiddleName = &rec[MiddleName]
		p.LastName = rec[LastName]
		p.Suffix = &rec[Suffix]
		p.Married = rec[MaritalStatus] == "Married"
		p.HomePhone = cleanPhoneNumber(rec[HomePhone])
		p.CellPhone = cleanPhoneNumber(rec[CellPhone])
		p.EmailAddress = &rec[Email]
		p.AddressLine1 = &rec[AddressLine1]
		p.AddressLine2 = &rec[AddressLine2]
		p.ZipCode = &rec[ZipCode]
		p.City = &rec[City]
		p.Province = &rec[State]
		p.Member = member(rec[MemberStatus])
		p.Birthday = pg.NullTime{Time: *parseDate(rec[BirthDate])}
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
		conn.Insert(&p)

		// setup events
		if d := parseDate(rec[BaptizedDate]); d != nil {
			e := db.Event{EventTypeRefer: eventBaptism.ID, PersonRefer: p.ID, Date: *d}
			conn.Insert(&e)
		}
		if d := parseDate(rec[DeceasedDate]); d != nil {
			e := db.Event{EventTypeRefer: eventDeath.ID, PersonRefer: p.ID, Date: *d}
			conn.Insert(&e)
		}
		if d := parseDate(rec[DateRemoved]); d != nil {
			e := db.Event{EventTypeRefer: eventLeftChurch.ID, PersonRefer: p.ID, Date: *d}
			conn.Insert(&e)
		}
		if d := parseDate(rec[WeddingDate]); d != nil {
			e := db.Event{EventTypeRefer: eventWedding.ID, PersonRefer: p.ID, Date: *d}
			conn.Insert(&e)
		}

		// set up background check
		if d := parseDate(rec[BackgroundCheck]); d != nil {
			e := db.Clearance{ClearanceTypeRefer: int64(clearanceBackgroundCheck.ID), PersonRefer: int64(p.ID), Date: pg.NullTime{Time: *d}}
			conn.Insert(&e)
		}

		// setup roles
		roles := strings.Split(rec[Leadership], "; ")
		for _, v := range roles {
			r := db.Role{Name: v}
			_, err = conn.Model(r).
				Column("id").
				Where("name = ?name").
				Returning("id").
				SelectOrInsert()
			if err != nil {
				return err
			}
			ra := db.RoleAssignment{PersonRefer: p.ID, RoleRefer: r.ID}
			conn.Insert(&ra)
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
				fc := db.ParentChildRelationship{ParentRefer: f.father.ID, ChildRefer: c.ID}
				conn.Insert(&fc)
			}
			if f.mother != nil {
				mc := db.ParentChildRelationship{ParentRefer: f.mother.ID, ChildRefer: c.ID}
				conn.Insert(&mc)
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
