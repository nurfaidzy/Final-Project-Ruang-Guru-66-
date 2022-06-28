package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type KampusRepository struct {
	db *sql.DB
}

func NewkampusRepository(db *sql.DB) *KampusRepository {
	return &KampusRepository{db: db}
}

func (k *KampusRepository) GetListKampus() ([]Kampus, error) {
	var sqlStatement string
	var kampuss []Kampus

	sqlStatement = `
		SELECT 
			id_kampus,
			nama_kampus,
			alamat_kampus,
			kota_kampus,
			provinsi_kampus,
			kode_pos_kampus,
			telepon_kampus,
			email_kampus,
			website_kampus,
			logo_kampus,
			id_user
		FROM kampus
	`
	rows, err := k.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var kampus Kampus
		err := rows.Scan(
			&kampus.Id_kampus,
			&kampus.Nama_kampus,
			&kampus.Alamat_kampus,
			&kampus.Kota_kampus,
			&kampus.Provinsi_kampus,
			&kampus.Kode_pos_kampus,
			&kampus.Telepon_kampus,
			&kampus.Email_kampus,
			&kampus.Website_kampus,
			&kampus.Logo_Kampus,
			&kampus.Id_user,
		)
		if err != nil {
			return nil, err
		}
		kampuss = append(kampuss, kampus)
	}
	return kampuss, nil
}

func (k *KampusRepository) InsertKampus(Nama_kampus string, Alamat_kampus string, Kota_kampus string, Provinsi_kampus string, Kode_pos_kampus string, Telepon_kampus string, Email_kampus string, Website_kampus string, Logo_Kampus string) error {
	_, err := k.db.Exec("INSERT INTO kampus (nama_kampus, alamat_kampus, kota_kampus, provinsi_kampus, kode_pos_kampus, telepon_kampus, email_kampus, website_kampus, logo_kampus) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", Nama_kampus, Alamat_kampus, Kota_kampus, Provinsi_kampus, Kode_pos_kampus, Telepon_kampus, Email_kampus, Website_kampus, Logo_Kampus)
	if err != nil {
		return err
	}
	return nil
}
