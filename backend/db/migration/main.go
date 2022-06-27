package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../../TARKAM.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id integer not null primary key AUTOINCREMENT,
			username varchar(255) not null,
			email varchar(255) not null,
			password varchar(255) not null
		);
		CREATE TABLE IF NOT EXISTS kampus (
			id integer not null primary key AUTOINCREMENT,
			name varchar(255) not null,
			email varchar(255) not null,
			jurusan1 varchar(255) not null,
			jurusan2 varchar(255) not null
		);	
		CREATE TABLE IF NOT EXISTS jurusan (
			id integer not null primary key AUTOINCREMENT,
			name varchar(255) not null
		);
		CREATE TABLE IF NOT EXISTS review (
			id integer not null primary key AUTOINCREMENT,
			username varchar(255) not null,
			kampus_name varhar(255) not null,
			jurusan_name varchar(255) not null,
			isian varhar(255) not null
		);
		Insert into users (username, email, password) values ("user","user","user" );
		Insert into kampus (name, email, jurusan1, jurusan2) values ("Universitas Indonesia", "humas-ui@ui.ac.id", "Ilmu Komputer", "Pendidikan Dokter");
		Insert into jurusan (name) values ("Ilmu Komputer");
		Insert into review (username, kampus_name, jurusan_name, isian) values ("user", "Universitas Indonesia", "Ilmu Komputer", "isian");
	`)
	if err != nil {
		panic(err)
	}
}
