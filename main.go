package main

import (
	"github.com/MikhaelEmpi/task-5-vix-btpns-devimikhaelempi/database"
	"github.com/MikhaelEmpi/task-5-vix-btpns-devimikhaelempi/models"
	"github.com/MikhaelEmpi/task-5-vix-btpns-devimikhaelempi/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
