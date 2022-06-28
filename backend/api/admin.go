package api

import "github.com/rg-km/final-project-engineering-66/backend/repository"

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
	Logo_kampus     string `json:"logo_kampus"`
}

type AddJurusan struct {
	Nama_jurusan string `json:"nama_jurusan"`
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
