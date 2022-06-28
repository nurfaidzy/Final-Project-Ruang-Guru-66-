package api

import (
	"fmt"
	"net/http"

	"github.com/rg-km/final-project-engineering-66/backend/repository"
)

type API struct {
	usersRepo   repository.UserRepository
	kampusRepo  repository.KampusRepository
	jurusanRepo repository.JurusanRepository
	reviewRepo  repository.ReviewRepository
	mux         *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, kampusRepo repository.KampusRepository, jurusanRepo repository.JurusanRepository, reviewRepo repository.ReviewRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, kampusRepo, jurusanRepo, reviewRepo, mux,
	}
	mux.Handle("/api/user/login", http.HandlerFunc(api.login))
	mux.Handle("/api/user/register", http.HandlerFunc(api.register))
	mux.Handle("/api/user/logout", http.HandlerFunc(api.logout))
	mux.Handle("/api/kampus/list", http.HandlerFunc(api.Kampuslist))
	mux.Handle("/api/kampus/add", http.HandlerFunc(api.AddKampus))
	mux.Handle("/api/jurusan/list", http.HandlerFunc(api.JurusanList))
	mux.Handle("/api/jurusan/add", http.HandlerFunc(api.AddJurusan))
	mux.Handle("/api/review/insetreview", http.HandlerFunc(api.InsertReview))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.mux)

}
