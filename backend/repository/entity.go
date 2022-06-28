package repository

type User struct {
	ID       int64  `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Jurusan struct {
	Id_jurusan   int64  `json:"id_jurusan"`
	Nama_jurusan string `json:"nama_jurusan"`
	Id_user      int64  `json:"id_user"`
}

type Kampus struct {
	Id_kampus   int64  `json:"id_kampus"`
	Nama_kampus string `json:"nama_kampus"`
	Id_user     int64  `json:"id_user"`
	// &kampus.Alamat_kampus,
	// &kampus.Kota_kampus,
	// &kampus.Provinsi_kampus,
	// &kampus.Kode_pos_kampus,
	// &kampus.Telepon_kampus,
	// &kampus.Email_kampus,
	// &kampus.Website_kampus,
	// &kampus.Logo_kampus,
	Alamat_kampus   string `json:"alamat_kampus"`
	Kota_kampus     string `json:"kota_kampus"`
	Provinsi_kampus string `json:"provinsi_kampus"`
	Kode_pos_kampus string `json:"kode_pos_kampus"`
	Telepon_kampus  string `json:"telepon_kampus"`
	Email_kampus    string `json:"email_kampus"`
	Website_kampus  string `json:"website_kampus"`
	Logo_kampus     string `json:"logo_kampus"`
}

type Review struct {
	Id_review  int64  `json:"id_review"`
	Id_jurusan int64  `json:"id_jurusan"`
	Id_user    int64  `json:"id_user"`
	Rating     int64  `json:"rating"`
	Review     string `json:"review"`
}

func (r *Review) IsEmpetyReview() bool {
	if r.Rating == 0 && r.Review == "" {
		return true
	}
	return false
}

func (r *Review) IsValidReview() bool {
	if r.Rating > 0 && r.Review != "" {
		return true
	}
	return false
}
