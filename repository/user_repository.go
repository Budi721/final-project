package repository

import (
	"github.com/itp-backend/backend-a-co-create/common/errors"
	"github.com/itp-backend/backend-a-co-create/model/domain"
	errCheck "github.com/pkg/errors"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(namaLengkap, username, password, topik, loginAs string) (*domain.User, error)
	FindByUsername(username, password, loginAs string) (*domain.User, error)
	DoesUsernameExist(username string) (bool, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{DB: db}
}

type userRepository struct {
	DB *gorm.DB
}

func (repo *userRepository) Create(namaLengkap, username, password, topik, loginAs string) (*domain.User, error) {
	user := &domain.User{
		Username: username,
		LoginAs:  loginAs,
	}
	user.Password = user.SetPassword(password)

	tx := repo.DB.Begin()
	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	enrollment := &domain.Enrollment{
		IdUser:           user.Id,
		NamaLengkap:      namaLengkap,
		Username:         username,
		TopikDiminati:    topik,
		EnrollmentStatus: 0,
	}
	enrollment.Password = user.SetPassword(password)

	result = tx.Create(&enrollment)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return user, nil
}

func (repo *userRepository) FindByUsername(username, password, loginAs string) (*domain.User, error) {
	u := &domain.User{}
	result := repo.DB.Where("username = ? AND login_as = ?", username, loginAs).First(u)

	if err := u.ComparePassword(password); err != nil {
		return nil, errors.NewInternalError(err, "Error: not found")
	}

	switch result.Error {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.NewInternalError(result.Error, "Error: not found")
	default:
		return nil, errors.NewInternalError(result.Error, "Error: database error")
	}
}

func (repo *userRepository) DoesUsernameExist(username string) (bool, error) {
	u := &domain.User{}
	if err := repo.DB.Where("username = ?", username).First(u).Error; errCheck.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, nil
}
