package main

import (
	"basictrade/database"
	"basictrade/router"
)

var (
	PORT = ":7070"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
