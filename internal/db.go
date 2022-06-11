package internal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("reducer.db"), &gorm.Config{})
}

func InitDB() (*gorm.DB, error) {
	db, err := GetDBConnection()
	if err != nil {
		return db, err
	}
	db.AutoMigrate(&URL{})
	return db, err
}
