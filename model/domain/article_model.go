package domain

type Article struct {
	IdArtikel   int    `json:"id_artikel,omitempty" gorm:"primaryKey"`
	PostingDate int64    `json:"posting_date,omitempty"`
	Kategori    string `json:"kategori,omitempty"`
	Judul       string `json:"judul,omitempty"`
	IsiArtikel  string `json:"isi_artikel,omitempty"`
	IdUser      int    `json:"id_user,omitempty"`
}
