package main

import (
	"log"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitMongoDB()
	db.InitDB()
	InitRouter()
}
