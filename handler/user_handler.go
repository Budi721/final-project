package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/common/responder"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    "github.com/itp-backend/backend-a-co-create/service"
    log "github.com/sirupsen/logrus"
    "net/http"
)

func Register(IUserService service.IUserService) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method != http.MethodPost {
            responder.NewHttpResponse(c, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        var register dto.Register

        if err := c.ShouldBindJSON(&register); err != nil {
            log.Warning(err)
            responder.NewHttpResponse(c, http.StatusBadRequest, nil, err)
            return
        }

        u, err := IUserService.Register(register)
        if err != nil {
            log.Warning(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusCreated, u, nil)
    }
}


func Login(IUserService service.IUserService) gin.HandlerFunc {
   return func(c *gin.Context) {
       if c.Request.Method != http.MethodPost {
           responder.NewHttpResponse(c, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
           return
       }
       var login dto.Login
       if err := c.ShouldBindJSON(&login); err != nil {
           log.Error(err)
           responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
           return
       }

       u, err := IUserService.Login(login.Username, login.Password, login.LoginAs)
       if err != nil {
           responder.NewHttpResponse(c, http.StatusUnauthorized, nil, errors.NewUnauthorizedError("invalid username and password"))
           return
       }

       responder.NewHttpResponse(c, http.StatusOK, u, nil)
       return
   }
}















