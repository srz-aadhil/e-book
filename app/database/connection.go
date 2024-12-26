package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	user     = "postgres"
// 	password = "password"
// 	host     = "localhost"
// 	port     = 5432
// 	dbname   = "ebookdb"
// )

func Initdb() (*gorm.DB, *sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s", user, password, host, port, dbname)

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to create instance of db", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Database ping failed", err)
	}

	fmt.Println("Database connection successfully created")
	return db, sqlDB, nil
}
