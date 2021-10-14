package dto

type Register struct {
	NamaLengkap   string `json:"nama_lengkap" binding:"required"`
	Username      string `json:"username" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	LoginAs       string   `json:"login_as" binding:"required"`
	TopikDiminati string `json:"topik_diminati" binding:"required"`
}
