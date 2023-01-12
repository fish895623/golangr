package ma

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func InitDb(db *gorm.DB) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	db.AutoMigrate(&Product{})

	if err != nil {
		panic("Cannot connect to database")
	}

	return db
}
