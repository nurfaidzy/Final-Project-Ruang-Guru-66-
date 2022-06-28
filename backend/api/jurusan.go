package api

import (
	"encoding/json"
	"net/http"
)

type JurusanListErrorResponse struct {
	Error string `json:"error"`
}

type Jurusanjson struct {
	Id_jurusan   int64  `json:"id_jurusan"`
	Nama_kampus  string `json:"nama_kampus"`
	Nama_jurusan string `json:"nama_jurusan"`
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

	for _, jurus := range jurusans {
		response.Jurusan = append(response.Jurusan, Jurusanjson{
			Id_jurusan:   jurus.Id_jurusan,
			Nama_kampus:  jurus.Nama_kampus,
			Nama_jurusan: jurus.Nama_jurusan,
		})
	}

	encoder.Encode(response)
}
