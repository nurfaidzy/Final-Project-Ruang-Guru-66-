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
	var juruss []Jurusan
	sqlStatement = `
		SELECT 
			id_jurusan,
			nama_jurusan,
			id_user
		FROM jurusan
		
	`
	rows, err := j.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var jurus Jurusan
		err := rows.Scan(
			&jurus.Id_jurusan,
			&jurus.Nama_jurusan,
			&jurus.Id_user,
		)
		if err != nil {
			return nil, err
		}
		juruss = append(juruss, jurus)
	}
	return juruss, nil
}
