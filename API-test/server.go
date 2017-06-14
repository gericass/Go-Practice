package main

import (
	"net/http"

	"./models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

var db = models.DbConnect() //DBと接続

func setDb(c echo.Context) error {
	db.CreateTable(&models.User{}) //テーブルの生成

	return c.String(http.StatusOK, "DB has set!")
}

func putDb(c echo.Context) error {
	name := c.Param("name")
	user := models.User{Age: 18, Name: name, Num: 50} //INSERTするデータ
	db.Create(&user)                                  //INSERT
	return c.String(http.StatusOK, "Put record!")
}

func updateDb(c echo.Context) error { //データの更新
	user := models.User{}
	db.Where("name = ?", "geri").First(&user)
	user.Age = 500
	db.Save(&user)
	return c.String(http.StatusOK, "Updated!")
}

func multiUpdate(c echo.Context) error {
	user := []models.User{}
	db.Find(&user, "age = ?", 18)
	text := ""
	for _, v := range user {
		text = text + v.Name
	}
	return c.String(http.StatusOK, text)
}

func main() {
	e := echo.New()
	defer db.Close() //return直前にDBとの接続を解除

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/set", setDb)
	e.GET("/put/:name", putDb)
	e.GET("/update", updateDb)
	e.GET("/multi", multiUpdate)

	e.Logger.Fatal(e.Start(":1323"))
}
