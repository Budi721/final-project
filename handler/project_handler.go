package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/common/helper"
    "github.com/itp-backend/backend-a-co-create/common/responder"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    "github.com/itp-backend/backend-a-co-create/service"
    log "github.com/sirupsen/logrus"
    "net/http"
)


func CreateProjectHandler(service service.IProjectService) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method != http.MethodPost {
            responder.NewHttpResponse(c, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        userEmail, err := helper.GetEmailUserLogin(c)
        if err != nil && userEmail == "" {
            responder.NewHttpResponse(c, http.StatusBadRequest, nil, err)
            return
        }
        userLogin := helper.GetUserData(userEmail)

        var project dto.Project
        project.Admin = userLogin.Id
        project.CollaboratorUserId = []int{userLogin.Id}
        if err := c.ShouldBindJSON(&project); err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        createdProject, err := service.CreateProject(&project)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusCreated, createdProject, nil)
        return
    }
}

func DetailProjectHandler(service service.IProjectService) gin.HandlerFunc {
    return func(c *gin.Context) {

    }
}

func DeleteProjectHandler(service service.IProjectService) gin.HandlerFunc {
    return func(c *gin.Context) {

    }
}

func ProjectByInvitedUserIdHandler(service service.IProjectService) gin.HandlerFunc {
    return func(c *gin.Context) {

    }
}
