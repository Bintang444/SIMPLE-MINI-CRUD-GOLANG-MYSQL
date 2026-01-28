# Mini CRUD Golang + Login Multi User

Project ini adalah aplikasi **CRUD sederhana menggunakan Golang (net/http)** dengan fitur:
- Login multi user (admin / user)
- Manajemen data anggota
- Struktur project rapi (MVC sederhana)
- Menggunakan MySQL sebagai database

Project ini dibuat untuk **latihan backend Golang**.

---

## 🚀 Fitur

- 🔐 Login & Logout (menggunakan cookie)
- 👤 Multi user (role: admin & user)
- 📋 CRUD data anggota
  - Create
  - Read
  - Update
  - Delete
- 🎨 Support CSS & JS (static file)
- 🧱 Struktur kode terpisah (controller, routes, database)

---

## 🗂️ Struktur Folder

```
mini-crud/
│
├── controller/
│   ├── anggota.go
│   └── auth.go
│
├── database/
│   └── db.go
│
├── routes/
│   └── routes.go
│
├── views/
│   ├── index.html
│   ├── create.html
│   ├── update.html
│   └── login.html
│
├── static/
│   ├── css/
│   └── js/
│
├── main.go
└── README.md
```

---

## 🛠️ Teknologi yang Digunakan

- Golang
- net/http
- MySQL
- HTML Template (html/template)
- XAMPP / MySQL Server

---

## 🧩 Database

### Nama Database
```sql
ujikom_go
```

### Tabel `anggota`
```sql
CREATE TABLE anggota (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nama VARCHAR(100),
  hobi VARCHAR(100),
  alamat TEXT
);
```

### Tabel `users`
```sql
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50),
  password VARCHAR(50),
  role VARCHAR(20)
);
```

### Contoh Data User
```sql
INSERT INTO users (username, password, role)
VALUES ('admin', 'admin123', 'admin'),
('user', 'user123', 'user');
```

> ⚠️ Password masih **plain text** (untuk kebutuhan pembelajaran).

---

## ⚙️ Konfigurasi Database

File:
```
database/db.go
```

```go
dsn := "root:@tcp(localhost:3306)/your_database?parseTime=true"
```

Silakan sesuaikan:
- username
- password
- nama database

---

## ▶️ Cara Menjalankan Project

1. Pastikan MySQL aktif
2. Buat database & tabel
3. Jalankan perintah:
```bash
go run main.go
```
4. Buka browser:
```
http://localhost:8080
```

---

## 🔐 Akun Login

| Username | Password  | Role  |
|--------|-----------|-------|
| admin  | admin123  | admin |
| user   | user123   | user  |

---

## 📌 Catatan

- Project ini **belum menggunakan framework** (pure Go)
- Cocok untuk:
  - Latihan backend
  - Pemahaman dasar MVC Golang

---

## 👨‍💻 Author

* Nama: Bintang Eka Wardhana Syarifudin

* Kelas: XII RPL 2

Dibuat sebagai latihan & pembelajaran backend Golang.  
Feel free to fork & develop 🚀