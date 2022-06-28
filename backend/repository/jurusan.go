package repository

import "database/sql"

type JurusanRepository struct {
	db *sql.DB
}

func NewJurusanRepository(db *sql.DB) *JurusanRepository {
	return &JurusanRepository{db: db}
}

func (j *JurusanRepository) GetListJurusan() ([]Jurusan, error) {
	var sqlStatement string
	var jurusans []Jurusan
	sqlStatement = `
		SELECT 
			id_jurusan,
			nama_kampus,
			nama_jurusan
		FROM jurusan
	`
	rows, err := j.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var jurusan Jurusan
		err := rows.Scan(
			&jurusan.Id_jurusan,
			&jurusan.Nama_kampus,
			&jurusan.Nama_jurusan,
		)
		if err != nil {
			return nil, err
		}
		jurusans = append(jurusans, jurusan)
	}
	return jurusans, nil
}

func (j *JurusanRepository) InsertJurusan(Nama_kampus string, Nama_jurusan string) error {
	_, err := j.db.Exec("INSERT INTO jurusan (nama_kampus, nama_jurusan) VALUES (?, ?)", Nama_kampus, Nama_jurusan)
	if err != nil {
		return err
	}
	return nil
}
