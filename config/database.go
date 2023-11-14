package config

import (
	"belajar-go-echo/app/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading environment", err)
	}
	// look detail
	log.Println("Environment varibale loading sukses")
}

func GetDsn() string {
	LoadEnv()

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
}

func InitDB() *gorm.DB {
	dsn := GetDsn()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	InitMigrateDB()

	return DB
}

func InitMigrateDB() {
	DB.AutoMigrate(&model.User{})
}
