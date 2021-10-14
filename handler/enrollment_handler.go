package handler

import (
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/common/helper"
    "github.com/itp-backend/backend-a-co-create/common/responder"
    "github.com/itp-backend/backend-a-co-create/service"
    "net/http"
)

func GetEnrollmentByStatus(service service.IEnrollmentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")

		userEmail, err := helper.GetEmailUserLogin(c)
		if err != nil && userEmail == "" {
			responder.NewHttpResponse(c, http.StatusBadRequest, nil, err)
			return
		}

		loginAs := helper.GetLoginAs(userEmail)
		if loginAs != "2" {
			responder.NewHttpResponse(c, http.StatusForbidden, nil, errors.New("you're not admin"))
			return
		}

		enrollments, err := service.GetEnrollmentByStatus(status)
		if len(enrollments) == 0 {
			responder.NewHttpResponse(c, http.StatusNotFound, nil, errors.New("not found"))
			return
		}

		if err != nil {
			responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
			return
		}

		responder.NewHttpResponse(c, http.StatusOK, enrollments, nil)
	}
}

func ApproveEnrollment(service service.IEnrollmentService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var idUserEnrollments map[string][]uint

        userEmail, err := helper.GetEmailUserLogin(c)
        if err != nil && userEmail == "" {
            responder.NewHttpResponse(c, http.StatusBadRequest, nil, err)
            return
        }
        loginAs := helper.GetLoginAs(userEmail)
        if loginAs != "2" {
            responder.NewHttpResponse(c, http.StatusForbidden, nil, errors.New("you're not admin"))
            return
        }

        if err := c.ShouldBindJSON(&idUserEnrollments); err != nil {
            responder.NewHttpResponse(c, http.StatusBadRequest, nil, err)
            return
        }

        enrollments, err := service.ApproveEnrollment(idUserEnrollments["user_ids"])
        if len(enrollments) == 0 {
            responder.NewHttpResponse(c, http.StatusNotFound, nil, errors.New("not found"))
            return
        }

        if err != nil {
            responder.NewHttpResponse(c, http.StatusInternalServerError, nil, err)
            return
        }
        responder.NewHttpResponse(c, http.StatusOK, enrollments, nil)
    }
}
