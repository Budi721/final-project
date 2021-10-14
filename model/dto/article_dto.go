package dto

type Article struct {
	Kategori   string `json:"kategori" binding:"required"`
	Judul      string `json:"judul" binding:"required"`
	IsiArtikel string `json:"isi_artikel" binding:"required"`
	IdUser     int    `json:"id_user" binding:"required"`
}
