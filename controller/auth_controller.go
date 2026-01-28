package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")

			var role string
			err := db.QueryRow(
				"SELECT role FROM users WHERE username=? AND password=?",
				username, password,
			).Scan(&role)

			if err != nil {
				http.Redirect(w, r, "/login?error=1", http.StatusSeeOther)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "login",
				Value: role,
				Path:  "/",
			})

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		tmpl := template.Must(template.ParseFiles(filepath.Join("views", "login.html")))
		tmpl.Execute(w, nil)
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "login",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
