package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/me/artisan_full/config"
	"github.com/me/artisan_full/handlers"
	"github.com/me/artisan_full/models"
)

func init() {
	// load values from .env file into the system.
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {

	// get configuration values
	conf := config.New()

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())  // logger middleware will “wrap” recovery
	e.Use(middleware.Recover()) // as it is enumerated before in the Use calls
	e.Use(middleware.CORS())

	// add db to context
	gormParamteres := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		conf.PostgresDB.DbHost,
		conf.PostgresDB.DbPort,
		conf.PostgresDB.DbName,
		conf.PostgresDB.DbUsername,
		conf.PostgresDB.DbPassword)
	gormDB, err := gorm.Open("postgres", gormParamteres)
	if err != nil {
		panic("failed to connect to db")
	}
	config.GormDB = gormDB
	// Migrate the tables
	// models.GormDB.AutoMigrate(&models.Test{})
	config.GormDB.AutoMigrate(&models.Help{})

	defer config.GormDB.Close()

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
	port := fmt.Sprintf(":%s", conf.ServerConf.ServerPort)
	e.Logger.Fatal(e.Start(port))
}
