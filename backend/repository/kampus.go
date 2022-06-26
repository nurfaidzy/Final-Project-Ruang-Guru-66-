package repository

import "database/sql"

type KampusRepository struct {
	db *sql.DB
}

type Kampus struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Jurusan1 string `json:"jurusan1"`
	Jurusan2 string `json:"jurusan2"`
}

func NewKampusRepository(db *sql.DB) *KampusRepository {
	return &KampusRepository{db: db}
}

func (k *KampusRepository) FetchKampusByID(id int64) ([]*Kampus, error) {
	var kampus []*Kampus
	query := `
		SELECT 	id, name, email,jurusan1, jurusan2 FROM kampus WHERE id = ?
	`
	rows, err := k.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.Id, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return nil, err
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus, nil
}

func (k *KampusRepository) FindKampus(name string, email string, jurusan1 string, jurusan2 string) error {
	_, err := k.db.Exec("SELECT * FROM kampus WHERE name = ? AND email = ? AND jurusan1 = ? AND jurusan2 = ?", name, email, jurusan1, jurusan2)
	if err != nil {
		return err
	}
	return nil
}

func (k *KampusRepository) FetchKampusByName(name string) ([]*Kampus, error) {
	var kampus []*Kampus
	query := `
		SELECT 	id, name, email,jurusan1, jurusan2 FROM kampus WHERE name = ?
	`
	rows, err := k.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.Id, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return nil, err
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus, nil
}

func (k *KampusRepository) FecthAllKampus() (name string, email string, jurusan1 string, jurusan2 string) {
	var kampus []*Kampus
	query := `
		SELECT 	id, name, email,jurusan1, jurusan2 FROM kampus
	`
	rows, err := k.db.Query(query)
	if err != nil {
		return "", "", "", ""
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.Id, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return "", "", "", ""
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus[0].Name, kampus[0].Email, kampus[0].Jurusan1, kampus[0].Jurusan2
}

func (k *KampusRepository) InsertKampus(Name string, Email string, Jurusan1 string, Jurusan2 string) error {

	_, err := k.db.Exec("INSERT INTO kampus (name, email, jurusan1, jurusan2) VALUES (?, ?, ?, ?)", Name, Email, Jurusan1, Jurusan2)
	if err != nil {
		return err
	}
	return nil

}

// func (k *KampusRepository) InsertKampus() ([]*Kampus, error) {
// 	var kampus []*Kampus
// 	query := `
// 		SELECT 	id, name, email,jurusan1, jurusan2  FROM kampus
// 	`
// 	rows, err := k.db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var kampusTemp Kampus
// 		err := rows.Scan(&kampusTemp.Id, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
// 		if err != nil {
// 			return nil, err
// 		}
// 		kampus = append(kampus, &kampusTemp)
// 	}
// 	return kampus, nil
// }
