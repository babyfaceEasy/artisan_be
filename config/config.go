package config

import (
	"os"
	"strconv"
	"strings"
	"github.com/jinzhu/gorm"
)

var GormDB *gorm.DB

// DBConfig - holds the config needed to connect to a database.
type DBConfig struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
}

// ServerConfig - holds the config that relates to the server.
type ServerConfig struct {
	ServerPort string
}

// Config holds all the configs for this app.
type Config struct {
	PostgresDB DBConfig
	ServerConf ServerConfig
	DebugMode  bool
}

// New returns a new config struct
func New() *Config {
	return &Config{
		PostgresDB: DBConfig{
			DbHost:     getEnv("DBHOST", ""),
			DbPort:     getEnv("DBPORT", ""),
			DbName:     getEnv("DBNAME", ""),
			DbUsername: getEnv("DBUSERNAME", ""),
			DbPassword: getEnv("DBPASSWORD", ""),
		},
		ServerConf: ServerConfig{
			ServerPort: getEnv("SERVERPORT", "1323"),
		},
		DebugMode: getEnvAsBool("DEBUG_MODE", true),
	}
}

// simple func to read enviroment variable and return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// simple function to read an enviroment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err != nil {
		return value
	}

	return defaultVal
}

// helper to read an enviroment variable into a bool or return a default
func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv("name", "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// helper to read an enviroment variable into string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valueStr := getEnv(name, "")

	if valueStr == "" {
		return defaultVal
	}

	val := strings.Split(valueStr, sep)
	return val
}
