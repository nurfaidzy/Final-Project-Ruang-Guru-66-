package api

import (
	"encoding/json"
	"net/http"

	"github.com/rg-km/final-project-engineering-66/backend/repository"
)

type AdminErrorResponse struct {
	Error string `json:"error"`
}

type AddKampus struct {
	Nama_kampus     string `json:"nama_kampus"`
	Alamat_kampus   string `json:"alamat_kampus"`
	Kota_kampus     string `json:"kota_kampus"`
	Provinsi_kampus string `json:"provinsi_kampus"`
	Kode_pos_kampus string `json:"kode_pos_kampus"`
	Telepon_kampus  string `json:"telepon_kampus"`
	Email_kampus    string `json:"email_kampus"`
	Website_kampus  string `json:"website_kampus"`
	Logo_Kampus     string `json:"logo_kampus"`
}

type AddKampusSuccessResponse struct {
	Message string `json:"message"`
}

type AddJurusan struct {
	Nama_kampus  string `json:"nama_kampus"`
	Nama_jurusan string `json:"nama_jurusan"`
}

type AdddjurusanSuccessResponse struct {
	Message string `json:"message"`
}

type AddUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AdminResponse struct {
	Kampus  []repository.Kampus  `json:"kampus"`
	Jurusan []repository.Jurusan `json:"jurusan"`
	User    []repository.User    `json:"user"`
}

func (api *API) AddKampus(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var request AddKampus
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(AdminErrorResponse{Error: "Invalid request"})
		return
	}
	err = api.kampusRepo.InsertKampus(request.Nama_kampus, request.Alamat_kampus, request.Kota_kampus, request.Provinsi_kampus, request.Kode_pos_kampus, request.Telepon_kampus, request.Email_kampus, request.Website_kampus, request.Logo_Kampus)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(AdminErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(AddKampusSuccessResponse{Message: "Kampus berhasil ditambahkan"})
}

func (api *API) AddJurusan(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var request AddJurusan
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(AdminErrorResponse{Error: "Invalid request"})
		return
	}
	err = api.jurusanRepo.InsertJurusan(request.Nama_kampus, request.Nama_jurusan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(AdminErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(AdddjurusanSuccessResponse{Message: "Jurusan berhasil ditambahkan"})
}

func (api *AdminResponse) IsEmpty() bool {
	if len(api.Kampus) == 0 && len(api.Jurusan) == 0 && len(api.User) == 0 {
		return true
	}
	return false
}

func (api *AdminResponse) IsValid() bool {
	if len(api.Kampus) > 0 && len(api.Jurusan) > 0 && len(api.User) > 0 {
		return true
	}
	return false
}
