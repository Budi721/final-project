package service

import (
    "context"
    "github.com/itp-backend/backend-a-co-create/model/domain"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    "github.com/itp-backend/backend-a-co-create/repository"
    log "github.com/sirupsen/logrus"
)

type IProjectService interface {
    CreateProject(project *dto.Project) (*domain.Project, error)
    GetDetailProject(ctx context.Context, projectId int) (*domain.Project, error)
    DeleteProject(ctx context.Context, projectId int)
    // Terima undangan
    GetProjectByInvitedUser(ctx context.Context, invitedId int) (*domain.Project, error)
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

func (service projectService) GetDetailProject(ctx context.Context, projectId int) (*domain.Project, error) {
    panic("implement me")
}

func (service projectService) DeleteProject(ctx context.Context, projectId int) {
    panic("implement me")
}

func (service projectService) GetProjectByInvitedUser(ctx context.Context, invitedId int) (*domain.Project, error) {
    panic("implement me")
}
