package dto

type Login struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    LoginAs  string `json:"login_as" binding:"required"`
}
