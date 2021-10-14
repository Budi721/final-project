package service

import (
    "github.com/google/uuid"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/config"
    "github.com/itp-backend/backend-a-co-create/contract"
    "github.com/itp-backend/backend-a-co-create/external/jwt_client"
    "github.com/itp-backend/backend-a-co-create/model/domain"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    "github.com/itp-backend/backend-a-co-create/repository"
    "time"
)

type IUserService interface {
    Register(registerResource dto.Register) (*domain.User, error)
    Login(username, password, loginAs string) (*domain.User, error)
}

func NewUserService(repo repository.IUserRepository, appConfig *config.Config, jwtClient jwt_client.JWTClientInterface) IUserService {
    return &userService{
        r:         repo,
        appConfig: appConfig,
        jwtClient: jwtClient,
    }
}

type userService struct {
    r repository.IUserRepository
    appConfig *config.Config
    jwtClient jwt_client.JWTClientInterface
}

func (service userService) Register(register dto.Register) (*domain.User, error) {
    exists, err := service.r.DoesUsernameExist(register.Username)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("User already exists")
    }

    atClaims := contract.JWTMapClaim{
        Authorized: true,
        RequestID:  uuid.New().String(),
    }

    atClaims.Subject = register.Username
    atClaims.ExpiresAt =  time.Now().Add(time.Hour * 24).Unix()

    token, err := service.jwtClient.GenerateTokenStringWithClaims(atClaims, service.appConfig.JWTSecret)
    if err != nil {
        return &domain.User{}, errors.NewBadRequestError(err)
    }

    userSaved, err := service.r.Create(
        register.NamaLengkap,
        register.Username,
        register.Password,
        register.TopikDiminati,
        register.LoginAs,
    )
    userSaved.AuthToken = token
    return userSaved, err
}

func (service userService) Login( username, password, loginAs string) (*domain.User, error) {
    user, err := service.r.FindByUsername(username, password, loginAs)
    if err != nil {
        return &domain.User{}, err
    }

    atClaims := contract.JWTMapClaim{
        Authorized: true,
        RequestID:  uuid.New().String(),
    }

    atClaims.Subject = user.Username
    atClaims.ExpiresAt =  time.Now().Add(time.Hour * 24).Unix()

    token, err := service.jwtClient.GenerateTokenStringWithClaims(atClaims, service.appConfig.JWTSecret)
    if err != nil {
        return &domain.User{}, errors.NewBadRequestError(err)
    }

    return &domain.User{Id: user.Id, AuthToken: token}, err
}

