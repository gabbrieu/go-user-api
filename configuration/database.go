package configuration

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Something wrong happens when loading the env file: %s", err)
	}

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	if err != nil {
		log.Fatalf("Something wrong happens when trying to convert the string port: %s", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", host, user, password, databaseName, port)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalf("Could not open the database: %s", err)
	}

	return db
}
