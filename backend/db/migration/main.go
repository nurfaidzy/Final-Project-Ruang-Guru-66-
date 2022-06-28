package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../../database.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id_user INTEGER PRIMARY key AUTOINCREMENT,
			username text(255) NOT NULL,
			password text(255) NOT NULL,
			role varchar(255) NOT NULL 
		);
		CREATE TABLE IF NOT EXISTS kampus (
			id_kampus INTEGER PRIMARY key AUTOINCREMENT,
			nama_kampus text(255) NOT NULL,
			alamat_kampus text(255) NOT NULL,
			kota_kampus text(255) NOT NULL,
			provinsi_kampus text(255) NOT NULL,
			kode_pos_kampus INTEGER NOT NULL,
			telepon_kampus INTEGER NOT NULL,
			email_kampus text(255) NOT NULL,
			website_kampus text(255) NOT NULL,
			logo_kampus text(255) NOT NULL,
			id_user INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS jurusan (
			id_jurusan INTEGER PRIMARY key AUTOINCREMENT,
			nama_jurusan text(255) NOT NULL,
			id_kampus INTEGER NOT NULL,
			FOREIGN KEY (id_kampus) REFERENCES kampus(id_kampus)

		);
		CREATE TABLE IF NOT EXISTS review (
			id_review INTEGER PRIMARY key AUTOINCREMENT,
			id_jurusan INTEGER NOT NULL,
			id_user INTEGER NOT NULL,
			review text(255) NOT NULL,
			FOREIGN KEY (id_jurusan) REFERENCES jurusan(id_jurusan),
			FOREIGN KEY (id_user) REFERENCES users(id_user)
		);
		INSERT INTO users (username, password, role) VALUES ("admin", "admin", "admin");

		`)
	if err != nil {
		panic(err)
	}
}
