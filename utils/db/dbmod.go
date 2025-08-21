package db

import (
	"fmt"
	"log"
	"os"

	figures "github.com/Jac-Bo/gofigure/utils/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {
	host := os.Getenv("HOST")
	port := "5433"
	dbName := "gormdb"
	dbUser := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	// Fmt.Sprintf formats the string accordingly
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
		//formats string for connection to database
	)

	DB, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	DB.AutoMigrate(figures.AFigure{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful")
	fmt.Println(DB)
}
