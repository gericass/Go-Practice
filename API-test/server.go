package main

import (
	"net/http"

	"./models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db := models.DbConnect()
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/set/", func(c echo.Context) error {
		db.CreateTable(&models.User{})
		//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Test{})
		return c.String(http.StatusOK, "DB has set!")
	})

	e.GET("/put/", func(c echo.Context) error {
		user := models.User{Age: 18, Name: "Jinzhu", Num: 50}
		db.Create(&user) // => returns `true` as primary key is blank
		return c.String(http.StatusOK, "Put record!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
