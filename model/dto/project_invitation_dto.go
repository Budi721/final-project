package dto

type ProjectInvitation struct {
	IdProject int `json:"id_project" binding:"required"`
	IdUser    int `json:"id_user" binding:"required"`
}
