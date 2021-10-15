package service

import (
    "github.com/itp-backend/backend-a-co-create/model/domain"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    "github.com/itp-backend/backend-a-co-create/repository"
    log "github.com/sirupsen/logrus"
)

type IProjectService interface {
    CreateProject(project *dto.Project) (*domain.Project, error)
    GetDetailProject(projectId int) (*domain.Project, error)
    DeleteProject(projectId int) error
    GetProjectByInvitedUser(invitedId int) ([]*domain.Project, error)
    UpdateInvitation(project dto.ProjectInvitation) (*domain.Project, error)
}

func NewProjectService(projectRepository repository.IProjectRepository) IProjectService {
    return projectService{
        repo:   projectRepository,
    }
}

type projectService struct {
    repo   repository.IProjectRepository
}

func (service projectService) CreateProject(project *dto.Project) (*domain.Project, error) {
    projectToCreate, err := service.repo.Create(project)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return projectToCreate, nil
}

func (service projectService) GetDetailProject(projectId int) (*domain.Project, error) {
    project, err := service.repo.FindById(projectId)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return project, nil
}

func (service projectService) DeleteProject(projectId int) error {
    err := service.repo.Delete(projectId)
    if err != nil {
        log.Error(err)
        return err
    }
    return nil
}

func (service projectService) GetProjectByInvitedUser(invitedId int) ([]*domain.Project, error) {
    project, err := service.repo.FindByInvitedUserId(invitedId)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return project, nil
}

func (service projectService) UpdateInvitation(project dto.ProjectInvitation) (*domain.Project, error) {
    projectUpdated, err := service.repo.UpdateInvitation(project)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return projectUpdated, nil
}