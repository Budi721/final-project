package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/common/responder"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    "github.com/itp-backend/backend-a-co-create/service"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
)

func CreateArticleHandler(service service.IArticleService) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method != http.MethodPost {
            responder.NewHttpResponse(c, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        var articleRequest *dto.Article
        if err := c.ShouldBindJSON(&articleRequest); err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        article, err := service.CreateArticle(articleRequest)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusCreated, article, nil)
        return
    }
}

func DeleteArticleHandler(service service.IArticleService) gin.HandlerFunc {
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

        err = service.DeleteArticle(id)
        result := map[string]int{
            "id_artikel_deleted": id,
        }

        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusNoContent, result, nil)
        return
    }
}

func GetArticleByIdHandler(service service.IArticleService) gin.HandlerFunc {
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

        article, err := service.GetArticleById(id)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusOK, article, nil)
        return
    }
}

func GetAllArticleHandler(service service.IArticleService) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method != http.MethodGet {
            responder.NewHttpResponse(c, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        articles, err := service.GetAllArticle()
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(c, http.StatusOK, articles, nil)
        return
    }
}

