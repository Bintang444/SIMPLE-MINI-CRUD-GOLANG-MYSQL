package main

import (
	"log"
	"net/http"

	"mini-crud/database"
	"mini-crud/routes"
)

func main() {
	db := database.InitDB()
	server := http.NewServeMux()

	routes.Register(server, db)

	log.Println("Server jalan di http://localhost:8080")
	http.ListenAndServe(":8080", server)
}