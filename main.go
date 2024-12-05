package main

import (
	"fmt"
	"log"

	"github.com/ArjunMalhotra07/internal/model"
	"github.com/ArjunMalhotra07/internal/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := os.Getenv("DSN")
	dsn := "root:Witcher_Arjun7@tcp(127.0.0.1:3306)/Library?charset=utf8mb4&parseTime=True&loc=Local"
	driver, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening DB : ", err.Error())
		return
	}
	if err := driver.AutoMigrate(&model.Book{}, &model.Author{}, &model.Borrow{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
		return
	}
	routes.OpenServer(driver)
}
