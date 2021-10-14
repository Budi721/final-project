package helper

import (
	"github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/app"
    "github.com/itp-backend/backend-a-co-create/model/domain"
    log "github.com/sirupsen/logrus"
)

func GetEmailUserLogin(c *gin.Context) (string, error) {
	userEmail := c.MustGet("userEmail").(string)
	defer func() {
		if r := recover(); r != nil {
			log.Warnln("Cannot get email user")
		}
	}()

	return userEmail, nil
}

func GetLoginAs(email string) string {
    var user *domain.User
    app.Init().DB.First(&user, "username = ?", email)
    return user.LoginAs
}
