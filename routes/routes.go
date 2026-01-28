package routes

import (
	"database/sql"
	"net/http"

	"mini-crud/controller"
)

func Register(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/login", controller.Login(db))
	server.HandleFunc("/logout", controller.Logout())

	server.HandleFunc("/", controller.Index(db))
	server.HandleFunc("/create", controller.Create(db))
	server.HandleFunc("/update", controller.Update(db))
	server.HandleFunc("/delete", controller.Delete(db))

	server.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}