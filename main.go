package main

import (
	"github.com/Ismadrade/api-book/database"
	server2 "github.com/Ismadrade/api-book/server"
)

func main() {
	database.StartDB()
	
	server := server2.NewServer()

	server.Run()
}
