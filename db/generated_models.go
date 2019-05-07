//lint:file-ignore U1000 ignore unused code, it's generated
package db

import (
	"github.com/go-pg/pg"
	"time"
)

var Columns = struct {
	ClearanceType struct {
		ID, Builtin, CreatedAt, DeletedAt, Duration, Name, UpdatedAt string
	}
	Clearance struct {
		ID, ClearanceTypeRefer, CreatedAt, Date, DeletedAt, PersonRefer, UpdatedAt string
	}
	EventType struct {
		ID, Builtin, CreatedAt, DeletedAt, Name, UpdatedAt string
	}
	Event struct {
		ID, CreatedAt, Date, DeletedAt, EventTypeRefer, PersonRefer, UpdatedAt string
	}
	GopgMigration struct {
		CreatedAt, ID, Version string
	}
	ParentChildRelationship struct {
		ID, ChildRefer, CreatedAt, DeletedAt, ParentRefer, UpdatedAt string
	}
	Person struct {
		ID, AddressLine1, AddressLine2, Birthday, CellPhone, City, CreatedAt, DeletedAt, EmailAddress, FirstName, HomePhone, Homebound, LastName, Married, Member, MiddleName, Notes, OutOfArea, PicturePath, Province, Sex, Suffix, UpdatedAt, ZipCode string
	}
	RoleAssignment struct {
		ID, CreatedAt, DeletedAt, EndDate, PersonRefer, RoleRefer, StartDate, UpdatedAt string
	}
	Role struct {
		ID, CreatedAt, DeletedAt, Name, UpdatedAt string
	}
	User struct {
		ID, EmailAddress, FirstName, LastName, PasswordHash string
	}
}{
	ClearanceType: struct {
		ID, Builtin, CreatedAt, DeletedAt, Duration, Name, UpdatedAt string
	}{
		ID:        "id",
		Builtin:   "builtin",
		CreatedAt: "created_at",
		DeletedAt: "deleted_at",
		Duration:  "duration",
		Name:      "name",
		UpdatedAt: "updated_at",
	},
	Clearance: struct {
		ID, ClearanceTypeRefer, CreatedAt, Date, DeletedAt, PersonRefer, UpdatedAt string
	}{
		ID:                 "id",
		ClearanceTypeRefer: "clearance_type_refer",
		CreatedAt:          "created_at",
		Date:               "date",
		DeletedAt:          "deleted_at",
		PersonRefer:        "person_refer",
		UpdatedAt:          "updated_at",
	},
	EventType: struct {
		ID, Builtin, CreatedAt, DeletedAt, Name, UpdatedAt string
	}{
		ID:        "id",
		Builtin:   "builtin",
		CreatedAt: "created_at",
		DeletedAt: "deleted_at",
		Name:      "name",
		UpdatedAt: "updated_at",
	},
	Event: struct {
		ID, CreatedAt, Date, DeletedAt, EventTypeRefer, PersonRefer, UpdatedAt string
	}{
		ID:             "id",
		CreatedAt:      "created_at",
		Date:           "date",
		DeletedAt:      "deleted_at",
		EventTypeRefer: "event_type_refer",
		PersonRefer:    "person_refer",
		UpdatedAt:      "updated_at",
	},
	GopgMigration: struct {
		CreatedAt, ID, Version string
	}{
		CreatedAt: "created_at",
		ID:        "id",
		Version:   "version",
	},
	ParentChildRelationship: struct {
		ID, ChildRefer, CreatedAt, DeletedAt, ParentRefer, UpdatedAt string
	}{
		ID:          "id",
		ChildRefer:  "child_refer",
		CreatedAt:   "created_at",
		DeletedAt:   "deleted_at",
		ParentRefer: "parent_refer",
		UpdatedAt:   "updated_at",
	},
	Person: struct {
		ID, AddressLine1, AddressLine2, Birthday, CellPhone, City, CreatedAt, DeletedAt, EmailAddress, FirstName, HomePhone, Homebound, LastName, Married, Member, MiddleName, Notes, OutOfArea, PicturePath, Province, Sex, Suffix, UpdatedAt, ZipCode string
	}{
		ID:           "id",
		AddressLine1: "address_line1",
		AddressLine2: "address_line2",
		Birthday:     "birthday",
		CellPhone:    "cell_phone",
		City:         "city",
		CreatedAt:    "created_at",
		DeletedAt:    "deleted_at",
		EmailAddress: "email_address",
		FirstName:    "first_name",
		HomePhone:    "home_phone",
		Homebound:    "homebound",
		LastName:     "last_name",
		Married:      "married",
		Member:       "member",
		MiddleName:   "middle_name",
		Notes:        "notes",
		OutOfArea:    "out_of_area",
		PicturePath:  "picture_path",
		Province:     "province",
		Sex:          "sex",
		Suffix:       "suffix",
		UpdatedAt:    "updated_at",
		ZipCode:      "zip_code",
	},
	RoleAssignment: struct {
		ID, CreatedAt, DeletedAt, EndDate, PersonRefer, RoleRefer, StartDate, UpdatedAt string
	}{
		ID:          "id",
		CreatedAt:   "created_at",
		DeletedAt:   "deleted_at",
		EndDate:     "end_date",
		PersonRefer: "person_refer",
		RoleRefer:   "role_refer",
		StartDate:   "start_date",
		UpdatedAt:   "updated_at",
	},
	Role: struct {
		ID, CreatedAt, DeletedAt, Name, UpdatedAt string
	}{
		ID:        "id",
		CreatedAt: "created_at",
		DeletedAt: "deleted_at",
		Name:      "name",
		UpdatedAt: "updated_at",
	},
	User: struct {
		ID, EmailAddress, FirstName, LastName, PasswordHash string
	}{
		ID:           "id",
		EmailAddress: "email_address",
		FirstName:    "first_name",
		LastName:     "last_name",
		PasswordHash: "password_hash",
	},
}

var Tables = struct {
	ClearanceType struct {
		Name, Alias string
	}
	Clearance struct {
		Name, Alias string
	}
	EventType struct {
		Name, Alias string
	}
	Event struct {
		Name, Alias string
	}
	GopgMigration struct {
		Name, Alias string
	}
	ParentChildRelationship struct {
		Name, Alias string
	}
	Person struct {
		Name, Alias string
	}
	RoleAssignment struct {
		Name, Alias string
	}
	Role struct {
		Name, Alias string
	}
	User struct {
		Name, Alias string
	}
}{
	ClearanceType: struct {
		Name, Alias string
	}{
		Name:  "clearance_types",
		Alias: "t",
	},
	Clearance: struct {
		Name, Alias string
	}{
		Name:  "clearances",
		Alias: "t",
	},
	EventType: struct {
		Name, Alias string
	}{
		Name:  "event_types",
		Alias: "t",
	},
	Event: struct {
		Name, Alias string
	}{
		Name:  "events",
		Alias: "t",
	},
	GopgMigration: struct {
		Name, Alias string
	}{
		Name:  "gopg_migrations",
		Alias: "t",
	},
	ParentChildRelationship: struct {
		Name, Alias string
	}{
		Name:  "parent_child_relationships",
		Alias: "t",
	},
	Person: struct {
		Name, Alias string
	}{
		Name:  "people",
		Alias: "t",
	},
	RoleAssignment: struct {
		Name, Alias string
	}{
		Name:  "role_assignments",
		Alias: "t",
	},
	Role: struct {
		Name, Alias string
	}{
		Name:  "roles",
		Alias: "t",
	},
	User: struct {
		Name, Alias string
	}{
		Name:  "users",
		Alias: "t",
	},
}

type ClearanceType struct {
	tableName struct{} `sql:"clearance_types,alias:t" pg:",discard_unknown_columns"`

	ID        int64       `sql:"id,pk"`
	Builtin   bool        `sql:"builtin,notnull"`
	CreatedAt pg.NullTime `sql:"created_at"`
	DeletedAt pg.NullTime `sql:"deleted_at"`
	Duration  *int64      `sql:"duration"`
	Name      string      `sql:"name,notnull"`
	UpdatedAt pg.NullTime `sql:"updated_at"`
}

type Clearance struct {
	tableName struct{} `sql:"clearances,alias:t" pg:",discard_unknown_columns"`

	ID                 int64       `sql:"id,pk"`
	ClearanceTypeRefer int64       `sql:"clearance_type_refer,notnull"`
	CreatedAt          pg.NullTime `sql:"created_at"`
	Date               pg.NullTime `sql:"date"`
	DeletedAt          pg.NullTime `sql:"deleted_at"`
	PersonRefer        int64       `sql:"person_refer,notnull"`
	UpdatedAt          pg.NullTime `sql:"updated_at"`
}

type EventType struct {
	tableName struct{} `sql:"event_types,alias:t" pg:",discard_unknown_columns"`

	ID        int64       `sql:"id,pk"`
	Builtin   bool        `sql:"builtin,notnull"`
	CreatedAt pg.NullTime `sql:"created_at"`
	DeletedAt pg.NullTime `sql:"deleted_at"`
	Name      string      `sql:"name,notnull"`
	UpdatedAt pg.NullTime `sql:"updated_at"`
}

type Event struct {
	tableName struct{} `sql:"events,alias:t" pg:",discard_unknown_columns"`

	ID             int64       `sql:"id,pk"`
	CreatedAt      pg.NullTime `sql:"created_at"`
	Date           time.Time   `sql:"date,notnull"`
	DeletedAt      pg.NullTime `sql:"deleted_at"`
	EventTypeRefer int64       `sql:"event_type_refer,notnull"`
	PersonRefer    int64       `sql:"person_refer,notnull"`
	UpdatedAt      pg.NullTime `sql:"updated_at"`
}

type GopgMigration struct {
	tableName struct{} `sql:"gopg_migrations,alias:t" pg:",discard_unknown_columns"`

	CreatedAt pg.NullTime `sql:"created_at"`
	ID        int         `sql:"id,notnull"`
	Version   *int64      `sql:"version"`
}

type ParentChildRelationship struct {
	tableName struct{} `sql:"parent_child_relationships,alias:t" pg:",discard_unknown_columns"`

	ID          int64       `sql:"id,pk"`
	ChildRefer  int64       `sql:"child_refer,notnull"`
	CreatedAt   pg.NullTime `sql:"created_at"`
	DeletedAt   pg.NullTime `sql:"deleted_at"`
	ParentRefer int64       `sql:"parent_refer,notnull"`
	UpdatedAt   pg.NullTime `sql:"updated_at"`
}

type Person struct {
	tableName struct{} `sql:"people,alias:t" pg:",discard_unknown_columns"`

	ID           int64       `sql:"id,pk"`
	AddressLine1 *string     `sql:"address_line1"`
	AddressLine2 *string     `sql:"address_line2"`
	Birthday     pg.NullTime `sql:"birthday"`
	CellPhone    *string     `sql:"cell_phone"`
	City         *string     `sql:"city"`
	CreatedAt    pg.NullTime `sql:"created_at"`
	DeletedAt    pg.NullTime `sql:"deleted_at"`
	EmailAddress *string     `sql:"email_address"`
	FirstName    string      `sql:"first_name,notnull"`
	HomePhone    *string     `sql:"home_phone"`
	Homebound    *bool       `sql:"homebound"`
	LastName     string      `sql:"last_name,notnull"`
	Married      bool        `sql:"married,notnull"`
	Member       bool        `sql:"member,notnull"`
	MiddleName   *string     `sql:"middle_name"`
	Notes        *string     `sql:"notes"`
	OutOfArea    *bool       `sql:"out_of_area"`
	PicturePath  *string     `sql:"picture_path"`
	Province     *string     `sql:"province"`
	Sex          int         `sql:"sex,notnull"`
	Suffix       *string     `sql:"suffix"`
	UpdatedAt    pg.NullTime `sql:"updated_at"`
	ZipCode      *string     `sql:"zip_code"`
}

type RoleAssignment struct {
	tableName struct{} `sql:"role_assignments,alias:t" pg:",discard_unknown_columns"`

	ID          int64       `sql:"id,pk"`
	CreatedAt   pg.NullTime `sql:"created_at"`
	DeletedAt   pg.NullTime `sql:"deleted_at"`
	EndDate     pg.NullTime `sql:"end_date"`
	PersonRefer int64       `sql:"person_refer,notnull"`
	RoleRefer   int64       `sql:"role_refer,notnull"`
	StartDate   pg.NullTime `sql:"start_date"`
	UpdatedAt   pg.NullTime `sql:"updated_at"`
}

type Role struct {
	tableName struct{} `sql:"roles,alias:t" pg:",discard_unknown_columns"`

	ID        int64       `sql:"id,pk"`
	CreatedAt pg.NullTime `sql:"created_at"`
	DeletedAt pg.NullTime `sql:"deleted_at"`
	Name      string      `sql:"name,notnull"`
	UpdatedAt pg.NullTime `sql:"updated_at"`
}

type User struct {
	tableName struct{} `sql:"users,alias:t" pg:",discard_unknown_columns"`

	ID           int64   `sql:"id,pk"`
	EmailAddress string  `sql:"email_address,notnull"`
	FirstName    string  `sql:"first_name,notnull"`
	LastName     string  `sql:"last_name,notnull"`
	PasswordHash *[]byte `sql:"password_hash"`
}
