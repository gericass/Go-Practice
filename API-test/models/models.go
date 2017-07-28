package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Age  int
	Name string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Num  int    `gorm:"AUTO_INCREMENT"`
}

func DbConnect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=Keita dbname=postgres sslmode=disable password=Keita9156")
	if err != nil {
		panic("だめぽ")
	}

	return db
}
