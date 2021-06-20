package main

import (
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
)

func main() {
	db.InitDB()
	InitRouter()
}
