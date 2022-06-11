package internal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("reducer.db"), &gorm.Config{})
}
