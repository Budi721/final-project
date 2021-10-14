package middleware

import (
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/common/responder"
    "github.com/itp-backend/backend-a-co-create/config"
    "github.com/itp-backend/backend-a-co-create/contract"
    "github.com/itp-backend/backend-a-co-create/external/jwt_client"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
    "strings"
)

type AuthValidate interface {
    EnsureLoggedIn() gin.HandlerFunc
    EnsureNotLoggedIn() gin.HandlerFunc
}

type authValidate struct {
    appConfig *config.Config
    jwtClient jwt_client.JWTClientInterface
}

func NewAuthValidate(appConfig *config.Config, jwtClient jwt_client.JWTClientInterface) AuthValidate {
    return &authValidate{
        appConfig: appConfig,
        jwtClient: jwtClient,
    }
}

func (a authValidate) EnsureLoggedIn() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.Request.Header.Get("Authorization")

        if authHeader == "" {
            responder.NewHttpResponse(c, http.StatusUnauthorized, nil, errors.New("no token found"))
            return
        }

        headerParts := strings.Split(authHeader, " ")
        if len(headerParts) != 2 {
            responder.NewHttpResponse(c, http.StatusUnauthorized, nil, errors.New("invalid auth header"))
            return
        }

        if headerParts[0] != "Bearer" {
            responder.NewHttpResponse(c, http.StatusUnauthorized, nil, errors.New("unauthorized - no bearer"))
            return
        }

        token := headerParts[1]
        claims := jwt.MapClaims{}

        err := a.jwtClient.ParseTokenWithClaims(token, claims, a.appConfig.JWTSecret)
        if err != nil {
            log.Errorln(err)
            responder.NewHttpResponse(c, http.StatusUnauthorized, nil, errors.New("invalid parse token with claims"))
            return
        }

        authorized := fmt.Sprintf("%v", claims["authorized"])
        requestID := fmt.Sprintf("%v", claims["requestID"])

        if authorized == "" || requestID == "" {
            responder.NewHttpResponse(c, http.StatusUnauthorized, nil, errors.New("invalid payload"))
            return
        }

        ok, err := strconv.ParseBool(authorized)

        if err != nil || !ok {
            log.Errorln(err)
            responder.NewHttpResponse(c, http.StatusUnauthorized, nil, errors.New("invalid payload"))
            return
        }

        resp := &contract.JWTMapClaim{
            Authorized:     claims["authorized"].(bool),
            RequestID:      claims["requestID"].(string),
            StandardClaims: jwt.StandardClaims{},
        }
        log.Println(resp)
        c.Set("email", claims["sub"].(string))
        c.Next()
    }

}

func (a authValidate) EnsureNotLoggedIn() gin.HandlerFunc {
    return func(c *gin.Context) {

    }
}



