package service

import (
    "github.com/itp-backend/backend-a-co-create/model/domain"
    "github.com/itp-backend/backend-a-co-create/repository"
    log "github.com/sirupsen/logrus"
)

type IEnrollmentService interface {
    GetEnrollmentByStatus(status string) ([]*domain.Enrollment, error)
    ApproveEnrollment(idUsers []uint) ([]*domain.Enrollment, error)
}

type enrollmentService struct {
    r repository.IEnrollmentRepository
}

func NewEnrollmentService(repository repository.IEnrollmentRepository) IEnrollmentService {
    return &enrollmentService{r: repository}
}

func (e enrollmentService) GetEnrollmentByStatus(status string) ([]*domain.Enrollment, error) {
    enrollments, err := e.r.FindAllByStatus(status)
    if err != nil {
        log.Error(err)
        return []*domain.Enrollment{}, err
    }

    return enrollments, nil
}

func (e enrollmentService) ApproveEnrollment(idUsers []uint) ([]*domain.Enrollment, error) {
    enrollments, err := e.r.UpdateStatusEnrollment(idUsers)
    if err != nil {
        log.Error(err)
        return []*domain.Enrollment{}, err
    }

    return enrollments, nil
}
