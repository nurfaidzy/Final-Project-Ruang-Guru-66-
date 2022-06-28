package api

import (
	"encoding/json"
	"net/http"
)

type KampusListErrorResponse struct {
	Error string `json:"error"`
}

type Kampusjson struct {
	Id_kampus       string `json:"id_kampus"`
	Nama_kampus     string `json:"nama_kampus"`
	Alamat_kampus   string `json:"alamat_kampus"`
	Kota_kampus     string `json:"kota_kampus"`
	Provinsi_kampus string `json:"provinsi_kampus"`
	Kode_pos_kampus string `json:"kode_pos_kampus"`
	Telepon_kampus  string `json:"telepon_kampus"`
	Email_kampus    string `json:"email_kampus"`
	Website_kampus  string `json:"website_kampus"`
	Logo_kampus     string `json:"logo_kampus"`
}

type KampusListResponse struct {
	Kampus []Kampusjson `json:"kampus"`
}

func (api *API) Kampuslist(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)
	response := KampusListResponse{}
	response.Kampus = make([]Kampusjson, 0)
	kampuss, err := api.kampusRepo.GetListKampus()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(KampusListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}

	for _, kakampuss := range kampuss {
		response.Kampus = append(response.Kampus, Kampusjson{
			Nama_kampus:     kakampuss.Nama_kampus,
			Alamat_kampus:   kakampuss.Alamat_kampus,
			Kota_kampus:     kakampuss.Kota_kampus,
			Provinsi_kampus: kakampuss.Provinsi_kampus,
			Kode_pos_kampus: kakampuss.Kode_pos_kampus,
			Telepon_kampus:  kakampuss.Telepon_kampus,
			Email_kampus:    kakampuss.Email_kampus,
			Website_kampus:  kakampuss.Website_kampus,
			Logo_kampus:     kakampuss.Logo_kampus,
		})
	}
	encoder.Encode(response)
}
