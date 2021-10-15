package repository

import (
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/model/domain"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type IProjectRepository interface {
    Create(project *dto.Project) (*domain.Project, error)
    FindById(idProject int) (*domain.Project, error)
    Delete(idProject int) error
    FindByInvitedUserId(invitedId int) ([]*domain.Project, error)
    UpdateInvitation(project dto.ProjectInvitation) (*domain.Project, error)
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
    return &projectRepository{DB: db}
}

type projectRepository struct {
    DB *gorm.DB
}

func (p projectRepository) Create(project *dto.Project) (*domain.Project, error) {
    var invitedUserId []domain.User

    if len(project.InvitedUserId) > 0 {
        p.DB.Find(&invitedUserId, project.InvitedUserId)
    }

    log.Errorln(project.TanggalMulai)
    projectToCreate := &domain.Project{
        KategoriProject:    project.KategoriProject,
        NamaProject:        project.NamaProject,
        TanggalMulai:       project.TanggalMulai,
        LinkTrello:         project.LinkTrello,
        DeskripsiProject:   project.DeskripsiProject,
        InvitedUserId:      project.InvitedUserId,
        Admin:              project.Admin,
        UsersInvited:       invitedUserId,
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

    if err := p.DB.Preload("UsersInvited").Preload("UsersCollaborator").First(&project).Error; err != nil {
        log.Error(err)
        return nil, err
    }

    for _, collaborator := range project.UsersCollaborator {
        project.CollaboratorUserId = append(project.CollaboratorUserId, collaborator.Id)
    }

    for _, invited := range project.UsersInvited {
        project.InvitedUserId = append(project.InvitedUserId, invited.Id)
    }
    return &project, nil
}

func (p projectRepository) Delete(idProject int) error {
    var project domain.Project
    project.IdProject = idProject
    p.DB.Model(&project).Association("UsersInvited")
    p.DB.Model(&project).Association("UsersCollaborator")

    if err := p.DB.First(&project).Error; err != nil {
        log.Error(err)
        return err
    }

    if err := p.DB.Delete(&project).Error; err != nil {
        log.Error(err)
        return errors.New("Cannot delete record")
    }

    return nil
}

func (p projectRepository) FindByInvitedUserId(invitedId int) ([]*domain.Project, error) {
    var projects []*domain.Project
    var user []domain.User
    p.DB.Find(&user, invitedId)
    if err := p.DB.Where(&domain.Project{UsersInvited: user}).Preload("UsersInvited").Preload("UsersCollaborator").Find(&projects).Error; err != nil {
        log.Error(err)
        return nil, err
    }

    for _, project := range projects {
        for _, collaborator := range project.UsersCollaborator {
            project.CollaboratorUserId = append(project.CollaboratorUserId, collaborator.Id)
        }

        for _, invited := range project.UsersInvited {
            project.InvitedUserId = append(project.InvitedUserId, invited.Id)
        }
    }

    return projects, nil
}

func (p projectRepository) UpdateInvitation(project dto.ProjectInvitation) (*domain.Project, error) {
    var projectUpdated domain.Project
    var invitedUserId []domain.User

    p.DB.Find(&invitedUserId, project.IdUser)
    projectUpdated.IdProject = project.IdProject

    p.DB.Model(&projectUpdated).Association("UsersInvited").Delete(&domain.User{Id: project.IdUser})
    p.DB.Model(&projectUpdated).Association("UsersCollaborator").Append(&invitedUserId)

    if err := p.DB.Preload("UsersInvited").Preload("UsersCollaborator").First(&projectUpdated).Error; err != nil {
        log.Error(err)
        return &domain.Project{}, err
    }

    for _, collaborator := range projectUpdated.UsersCollaborator {
        projectUpdated.CollaboratorUserId = append(projectUpdated.CollaboratorUserId, collaborator.Id)
    }

    for _, invited := range projectUpdated.UsersInvited {
        projectUpdated.InvitedUserId = append(projectUpdated.InvitedUserId, invited.Id)
    }
    return &projectUpdated, nil
}