package api

import (
	"encoding/json"
	"net/http"
)

type JurusanListErrorResponse struct {
	Error string `json:"error"`
}

type Jurusanjson struct {
	Id_jurusan   string `json:"id_jurusan"`
	Nama_jurusan string `json:"nama_jurusan"`
	Id_user      string `json:"id_user"`
}

type JurusanListSuccessResponse struct {
	Jurusan []Jurusanjson `json:"jurusan"`
}

func (api *API) JurusanList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)
	response := JurusanListSuccessResponse{}
	response.Jurusan = make([]Jurusanjson, 0)
	jurusans, err := api.jurusanRepo.GetListJurusan()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(JurusanListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}

	for _, kajurusans := range jurusans {
		response.Jurusan = append(response.Jurusan, Jurusanjson{
			Nama_jurusan: kajurusans.Nama_jurusan,
		})
	}
	encoder.Encode(response)
}
