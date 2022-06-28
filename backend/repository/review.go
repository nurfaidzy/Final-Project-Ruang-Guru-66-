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
			id_jurusan,
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
			&review.Id_jurusan,
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

func (r *ReviewRepository) InsertReview(Review string) error {
	_, err := r.db.Exec("INSERT INTO review (review) VALUES ( ?)", Review)
	if err != nil {
		return err
	}
	return nil
}
