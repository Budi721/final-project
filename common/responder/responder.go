package responder

import (
    "github.com/gin-gonic/gin"
    "net/http"

    "github.com/itp-backend/backend-a-co-create/common/errors"
    pkgErrors "github.com/pkg/errors"
    "github.com/sirupsen/logrus"
)

type Template struct {
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

func NewHttpResponse(c *gin.Context, httpCode int, result interface{}, err error) {
	if err != nil {
		Error(c, err, httpCode)
	} else {
		Success(c, result, httpCode)
	}
}

func Error(c *gin.Context, err error, httpCode int) {
	switch err := pkgErrors.Cause(err).(type) {
	case *errors.BadRequestError:
		badRequestError(c, err)
	case *errors.UnauthorizedError:
		unauthorizedError(c, err)
	default:
		GenericError(c, err, err.Error(), httpCode)
	}
}

func Success(c *gin.Context, successResponse interface{}, responseCode ...int) {

	t := Template{
		Status: http.StatusOK,
        Error:  nil,
        Result: successResponse,
	}

    c.JSON(http.StatusOK, t)
}

func GenericError(c *gin.Context, err error, errorResponse interface{}, responseCode int) {
	log := logrus.WithFields(logrus.Fields{
		"Method": c.Request.Method,
		"Host":   c.Request.Host,
		"Path":   c.Request.RequestURI,
	}).WithField("ResponseCode", responseCode)

	if responseCode < 500 {
		log.Warn(err.Error())
	} else {
		log.Error(err.Error())
	}

	t := Template{
		Status: responseCode,
		Result: nil,
		Error:  errorResponse,
	}

	if errorResponse != nil {
		c.AbortWithStatusJSON(responseCode, t)
	}
}

func badRequestError(c *gin.Context, err *errors.BadRequestError) {
	GenericError(c, err, err.Error(), http.StatusBadRequest)
}

func unauthorizedError(c *gin.Context, err *errors.UnauthorizedError) {
	GenericError(c, err, err.Error(), http.StatusUnauthorized)
}
