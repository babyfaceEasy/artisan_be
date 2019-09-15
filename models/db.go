package models

import (
	"github.com/jinzhu/gorm"
)

type Configuration struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
}

var GormDB *gorm.DB
var configuration Configuration

func init() {
	configuration = Configuration{
		DbHost:     "localhost",
		DbPort:     "5432",
		DbName:     "artisan_db",
		DbUsername: "michael.olakunle",
		DbPassword: "postgres",
	}
}

func GetConfig() Configuration {
	return configuration
}