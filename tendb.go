package tendb

import "github.com/jinzhu/gorm"

func OpenDatabase() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=localhost user=postgres dbname=tendb sslmode=disable password=postgres")
}
