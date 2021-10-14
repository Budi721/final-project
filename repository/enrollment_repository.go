package repository

import (
    "github.com/itp-backend/backend-a-co-create/model/domain"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type IEnrollmentRepository interface {
    FindAllByStatus(status string) ([]*domain.Enrollment, error)
    UpdateStatusEnrollment(idUsers []uint) ([]*domain.Enrollment, error)
}

func NewEnrollmentRepository(db *gorm.DB) IEnrollmentRepository {
    return &enrollmentRepository{DB: db}
}

type enrollmentRepository struct {
    DB *gorm.DB
}

func (e enrollmentRepository) FindAllByStatus(status string) ([]*domain.Enrollment, error) {
    var enrollments []*domain.Enrollment
    if err := e.DB.Where("enrollment_status = ?", status).Find(&enrollments).Error; err != nil {
        log.Error(err)
        return enrollments, err
    }

    return enrollments, nil
}

func (e enrollmentRepository) UpdateStatusEnrollment(idUsers []uint) ([]*domain.Enrollment, error) {
    var enrollments []*domain.Enrollment
    e.DB.Where("id_user IN ?", idUsers).Find(&enrollments)

    if err := e.DB.Table("enrollments").
        Where("id_user IN ?", idUsers).
        Updates(domain.Enrollment{EnrollmentStatus: 1}).
        Error; err != nil {
        log.Error(err)
        return enrollments, err
    }

    return enrollments, nil
}

