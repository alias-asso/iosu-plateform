package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSqlite(path string) (error, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return err, &gorm.DB{}
	}
	return nil, db
}
