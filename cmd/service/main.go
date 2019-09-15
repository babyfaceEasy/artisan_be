package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/me/artisan_full/handlers"
	"github.com/me/artisan_full/models"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())  // logger middleware will “wrap” recovery
	e.Use(middleware.Recover()) // as it is enumerated before in the Use calls

	// e.File("/", "static/index.html")

	// add db to context
	configuration := models.GetConfig()
	gormParamteres := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", configuration.DbHost, configuration.DbPort, configuration.DbName, configuration.DbUsername, configuration.DbPassword)
	gormDB, err := gorm.Open("postgres", gormParamteres)
	if err != nil {
		panic("failed to connect to db")
	}
	models.GormDB = gormDB
	// Migrate the tables
	// models.GormDB.AutoMigrate(&models.Test{})
	models.GormDB.AutoMigrate(&models.Help{})

	defer models.GormDB.Close()

	// in order to serve static assets
	e.Static("/static", "static")

	// routes
	// V1 routes
	v1 := e.Group("v1")
	v1.POST("/", handlers.CreateHelp)
	/*
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
	*/
	e.Logger.Fatal(e.Start(":1323"))
}
