package main

import (
	"https://github.com/MikhaelEmpi/task-5-vix-btpns-devimikhaelempi/database"
	"https://github.com/MikhaelEmpi/task-5-vix-btpns-devimikhaelempi/models"
	"https://github.com/MikhaelEmpi/task-5-vix-btpns-devimikhaelempi/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
