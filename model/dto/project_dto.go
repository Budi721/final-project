package dto

type Project struct {
	KategoriProject    string `json:"kategori_project" binding:"required"`
	NamaProject        string `json:"nama_project" binding:"required"`
	TanggalMulai       int64  `json:"tanggal_mulai" binding:"required"`
	LinkTrello         string `json:"link_trello" binding:"required"`
	DeskripsiProject   string `json:"deskripsi_project" binding:"required"`
	InvitedUserId      []int  `json:"invited_user_id" binding:"required"`
	Admin              int    `json:"admin"`
}
