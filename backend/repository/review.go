package repository

import "database/sql"

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetListReview() ([]Review, error) {
	var sqlStatement string
	var reviews []Review
	sqlStatement = `
		SELECT 
			id_review,
			id_user,
			nama_kampus,
			nama_jurusan,
			review,
			rating
		FROM review
	`
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var review Review
		err := rows.Scan(
			&review.Id_review,
			&review.Id_user,
			&review.Nama_kampus,
			&review.Nama_jurusan,
			&review.Review,
			&review.Rating,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (r *ReviewRepository) InsertReview(Id_review int64, Id_user int64, Nama_kampus string, Nama_jurusan string, Review string, Rating int64) error {
	_, err := r.db.Exec("INSERT INTO review (id_review, id_user, nama_kampus, nama_jurusan, review, rating) VALUES (?, ?, ?, ?, ?, ?)", Id_review, Id_user, Nama_kampus, Nama_jurusan, Review, Rating)
	if err != nil {
		return err
	}
	return nil

}
