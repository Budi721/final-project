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
    "strconv"
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
        if c.Request.Method != http.MethodGet {
            responder.NewHttpResponse(c, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        param := c.Param("id")
        id, err := strconv.Atoi(param)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }
        project, err := service.GetDetailProject(id)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusOK, project, nil)
    }
}

func DeleteProjectHandler(service service.IProjectService) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method != http.MethodDelete {
            responder.NewHttpResponse(c, http.StatusMethodNotAllowed, "method not allowed", errors.New("Error: Method is not allowed"))
            return
        }

        param := c.Param("id")
        id, err := strconv.Atoi(param)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        err = service.DeleteProject(id)
        result := map[string]int{
            "id_project_deleted": id,
        }

        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusNoContent, result, nil)
    }
}

func ProjectByInvitedUserIdHandler(service service.IProjectService) gin.HandlerFunc {
    return func(c *gin.Context) {
        invitedUser := c.Query("invited_user_id")
        id, err := strconv.Atoi(invitedUser)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        project, err := service.GetProjectByInvitedUser(id)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusOK, project, nil)
    }
}

func AcceptProjectHandler(service service.IProjectService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var projectInvitation dto.ProjectInvitation

        if err := c.ShouldBindJSON(&projectInvitation); err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        projectUpdated, err := service.UpdateInvitation(projectInvitation)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusOK, projectUpdated, nil)
        return
    }
}
