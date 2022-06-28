package api

import (
	"encoding/json"
	"net/http"
)

type ReviewsErrorResponse struct {
	Error string `json:"error"`
}

type ReviewReviews struct {
	Review string `json:"review"`
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
	err = api.reviewRepo.InsertReview(reviews.Review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ReviewsErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(ReviewsSuccess{Message: "Review berhasil ditambahkan"})

}
