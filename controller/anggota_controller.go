package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type Anggota struct {
	Id int
	Nama string
	Hobi string
	Alamat string
}

/* ================= INDEX ================= */
func Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nama, hobi, alamat FROM anggota")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		var data []Anggota
		for rows.Next() {
			var a Anggota
			rows.Scan(&a.Id, &a.Nama, &a.Hobi, &a.Alamat)
			data = append(data, a)
		}

		tmpl := template.Must(template.ParseFiles(filepath.Join("views", "index.html")))
		tmpl.Execute(w, data)
	}
}

/* ================= CREATE ================= */
func Create(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			nama := r.FormValue("nama")
			hobi := r.FormValue("hobi")
			alamat := r.FormValue("alamat")

			db.Exec("INSERT INTO anggota (nama, hobi, alamat) VALUES (?,?,?)",
				nama, hobi, alamat)

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		tmpl := template.Must(template.ParseFiles(filepath.Join("views", "create.html")))
		tmpl.Execute(w, nil)
	}
}

/* ================= UPDATE ================= */
func Update(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		if r.Method == http.MethodPost {
			nama := r.FormValue("nama")
			hobi := r.FormValue("hobi")
			alamat := r.FormValue("alamat")

			db.Exec("UPDATE anggota SET nama=?, hobi=?, alamat=? WHERE id=?",
				nama, hobi, alamat, id)

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		var a Anggota
		db.QueryRow("SELECT nama, hobi, alamat FROM anggota WHERE id=?", id).
			Scan(&a.Nama, &a.Hobi, &a.Alamat)
		a.Id, _ = strconv.Atoi(id)

		tmpl := template.Must(template.ParseFiles(filepath.Join("views", "update.html")))
		tmpl.Execute(w, a)
	}
}

/* ================= DELETE ================= */
func Delete(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		db.Exec("DELETE FROM anggota WHERE id=?", id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func isAdmin(r *http.Request) bool {
	cookie, err := r.Cookie("login")
	return err == nil && cookie.Value == "admin"
}