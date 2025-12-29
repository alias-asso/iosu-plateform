package database

import (
	"errors"

	"github.com/alias-asso/iosu/plateform/config"
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

func ConnectDb(config *config.Config) (error, *gorm.DB) {
	var db *gorm.DB
	switch config.DbType {
	case "sqlite":
		err, conn := ConnectSqlite(config.Sqlite.DbPath)
		if err != nil {
			return err, &gorm.DB{}
		}
		db = conn
	case "postgres":
		// TODO
		return errors.ErrUnsupported, &gorm.DB{}

	case "mysql":
		// TODO
		return errors.ErrUnsupported, &gorm.DB{}

	default:
		// TODO
		return errors.ErrUnsupported, &gorm.DB{}
	}
	return nil, db
}
