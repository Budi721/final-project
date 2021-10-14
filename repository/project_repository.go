package repository

import (
    "github.com/itp-backend/backend-a-co-create/model/domain"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type IProjectRepository interface {
    Create(project *dto.Project) (*domain.Project, error)
    FindById(idProject int) (*domain.Project, error)
    Delete(idProject int) error
    FindByInvitedUserId(invitedId int) (*domain.Project, error)
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
    return projectRepository{DB: db}
}

type projectRepository struct {
    DB *gorm.DB
}

func (p projectRepository) Create(project *dto.Project) (*domain.Project, error) {
    var invitedUserId []domain.User
    var collaboratorUserId []domain.User

    if len(project.InvitedUserId) > 0 {
        p.DB.Find(&invitedUserId, project.InvitedUserId)
    }
    if len(project.CollaboratorUserId) > 0 {
        p.DB.Find(&collaboratorUserId, project.CollaboratorUserId)
    }

    log.Errorln(project.TanggalMulai)
    projectToCreate := &domain.Project{
        KategoriProject:    project.KategoriProject,
        NamaProject:        project.NamaProject,
        TanggalMulai:       project.TanggalMulai,
        LinkTrello:         project.LinkTrello,
        DeskripsiProject:   project.DeskripsiProject,
        InvitedUserId:      project.InvitedUserId,
        CollaboratorUserId: project.CollaboratorUserId,
        Admin:              project.Admin,
        UsersInvited:       invitedUserId,
        UsersCollaborator:  collaboratorUserId,
    }
    result := p.DB.Create(&projectToCreate)
    if result.Error != nil {
        log.Error(result.Error)
        return nil, result.Error
    }
    return projectToCreate, nil
}

func (p projectRepository) FindById(idProject int) (*domain.Project, error) {
    var project domain.Project
    project.IdProject = idProject

    if err := p.DB.First(&project).Error; err != nil {
        log.Error(err)
        return nil, err
    }

    return &project, nil
}

func (p projectRepository) Delete(idProject int) error {
    var project domain.Project
    project.IdProject = idProject

    if err := p.DB.Delete(&project).Error; err != nil {
        log.Error(err)
        return err
    }

    return nil
}

func (p projectRepository) FindByInvitedUserId(invitedId int) (*domain.Project, error) {
    var project domain.Project
    var user []domain.User

    p.DB.Find(&user, invitedId)
    project.UsersInvited = user
    if err := p.DB.First(&project).Error; err != nil {
        log.Error(err)
        return nil, err
    }

    return &project, nil
}

