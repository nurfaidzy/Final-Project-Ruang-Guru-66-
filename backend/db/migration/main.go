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
			nama_kampus text(255) NOT NULL,
			nama_jurusan text(255) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS review (
			id_review INTEGER PRIMARY key AUTOINCREMENT,
			id_user INTEGER NOT NULL,
			nama_kampus text(255) NOT NULL,
			nama_jurusan text(255) NOT NULL,
			review text(255) NOT NULL,
			rating INTEGER NOT NULL
		);
		INSERT INTO users (username, password, role) VALUES ("admin", "admin", "admin");
		INSERT INTO kampus (nama_kampus, alamat_kampus, kota_kampus, provinsi_kampus, kode_pos_kampus, telepon_kampus, email_kampus, website_kampus, logo_kampus, id_user) 
			VALUES ("Universitas Indonesia", "jl. merdeka no.1", "Jakarta", "DKI Jakarta", "12345", "12345", "ui@.gmail.com", "ui.com", "logo.jpg", 1),
			("Universitas Padjadjaran", "jl.bandung", "Bandung", "Jawa Barat", "12345", "12345", "up@.gmail.com", "up.com", "logo.jpg", 2),
			("Institut Teknologi Bandung", "jl.bandung", "Bandung", "Jawa Barat", "12345", "12345", "itb@.gmail.com", "itb.com", "logo.jpg", 3),
			("Institut Teknologi Sepuluh Nopember", "jl.bandung", "Bandung", "Jawa Barat", "12345", "12345", "its@.gmail.com", "its.com", "logo.jpg", 4);
		INSERT INTO jurusan (nama_kampus, nama_jurusan) VALUES ("Universitas Brawijaya", "Teknik Informatika");
		INSERT INTO review (id_user, nama_kampus, nama_jurusan, review, rating) VALUES (1, "Universitas Brawijaya", "Teknik Informatika", "Sangat bagus", 5);
		`)
	if err != nil {
		panic(err)
	}
}
