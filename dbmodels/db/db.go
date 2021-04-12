package adminmodels

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDbConnection handles the connection to a gorm database, so that DB CONNECTION Environment Variables are also standarized
func GetDbConnection(dbtype string, logLevel logger.LogLevel, devMode bool) *gorm.DB {
	_ = godotenv.Load()
	var dbName string
	if devMode {
		if os.Getenv("POSTGRES_DATABASE") != "" {
			dbName = os.Getenv("POSTGRES_DATABASE")
		} else {
			dbName = "bitcou_test"
		}
	} else {
		dbName = "bitcou_products"
	}
	
	if dbtype == "mysql" {
		credentials := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
		db, err := gorm.Open(mysql.Open(credentials), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
		if err != nil {
			panic("unable to connect to mysql")
		}
		return db
	}
	if dbtype == "postgres" {
		var err error
		credentials := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("POSTGRES_IP"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), dbName)
		db, err := gorm.Open(postgres.Open(credentials), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
		if err != nil {
			panic("unable to connect to postgres")
		}
		return db
	}
	panic("no supported database was selected")
}
