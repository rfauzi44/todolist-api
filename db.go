package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rfauzi44/todolist-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get env
	MYSQL_USER := getEnv("MYSQL_USER", "root")
	MYSQL_PASSWORD := getEnv("MYSQL_PASSWORD", "")
	MYSQL_HOST := getEnv("MYSQL_HOST", "127.0.0.1")
	MYSQL_PORT := getEnv("MYSQL_PORT", "3306")
	MYSQL_DBNAME := getEnv("MYSQL_DBNAME", "dbname")

	conString := MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DBNAME + "?charset=utf8mb4&parseTime=True"
	fmt.Println(conString)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// Set connection pool options
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(1000)

	// Auto-migrate tables
	err = db.AutoMigrate(&model.Activity{}, &model.Todo{})
	if err != nil {
		panic(err)
	}

	return db
}
