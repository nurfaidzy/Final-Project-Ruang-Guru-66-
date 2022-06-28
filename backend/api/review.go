package api

import (
	"encoding/json"
	"net/http"
)

type ReviewsErrorResponse struct {
	Error string `json:"error"`
}

type ReviewReviews struct {
	Id_review    int64  `json:"id_review"`
	Id_user      int64  `json:"id_user"`
	Nama_kampus  string `json:"nama_kampus"`
	Nama_jurusan string `json:"nama_jurusan"`
	Review       string `json:"review"`
	Rating       int64  `json:"rating"`
}

type ReviewsSuccess struct {
	Message string `json:"message"`
}

func (api *API) InsertReview(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var reviews ReviewReviews
	err := json.NewDecoder(req.Body).Decode(&reviews)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ReviewsErrorResponse{Error: err.Error()})
		return
	}
	err = api.reviewRepo.InsertReview(reviews.Id_review, reviews.Id_user, reviews.Nama_kampus, reviews.Nama_jurusan, reviews.Review, reviews.Rating)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ReviewsErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(ReviewsSuccess{Message: "Review berhasil ditambahkan"})

}
