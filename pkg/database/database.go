package database

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = ConnectDB()
	if err != nil {
		return fmt.Errorf("error initializing PostgreSQL database: %v", err)
	}

	return nil
}

func GetPostgreSQL() *gorm.DB {
	return db
}
