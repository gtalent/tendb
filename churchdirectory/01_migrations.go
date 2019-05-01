//go:generate go run ../assets_generate/assets_generate.go churchdirectory
package churchdirectory

import (
	"io"
	"net/http"
	"strconv"

	"github.com/go-pg/migrations"
)

func ls(assets http.FileSystem) ([]string, error) {
	dir, err := assets.Open("/")
	if err != nil {
		return nil, err
	}
	list, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, f := range list {
		println(f.Name())
		out = append(out, f.Name())
	}
	return out, nil
}

func LsAssets() {
	ls(assets)
}

func loadMigrationFile(path string) (string, error) {
	file, err := assets.Open(path)
	if err != nil {
		return "", err
	}
	buff := make([]byte, 0)
	_, err = io.ReadFull(file, buff)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

func init() {
	loop := true
	for i := 0; loop; i++ {
		path := strconv.Itoa(i) + "_init.tx.up.sql"
		sql, err := loadMigrationFile(path)
		migrations.MustRegisterTx(func(db migrations.DB) error {
			if err != nil {
				loop = false
				return err
			}
			_, err = db.Exec(sql)
			return err
		}, func(db migrations.DB) error {
			path := strconv.Itoa(i) + ".tx.down.sql"
			sql, err := loadMigrationFile(path)
			if err != nil {
				loop = false
				return err
			}
			_, err = db.Exec(sql)
			return err
		})
	}
}
