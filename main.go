package main

import (
	"github.com/beruang43221/assignment2-015/database"
	"github.com/beruang43221/assignment2-015/routers"
)

func main() {
	database.StartDB()

	router := routers.SetupRouter()
	router.Run(":8083")
}