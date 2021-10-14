package service

import (
    "github.com/itp-backend/backend-a-co-create/model/domain"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    "github.com/itp-backend/backend-a-co-create/repository"
    log "github.com/sirupsen/logrus"
)

type IArticleService interface {
    CreateArticle(article *dto.Article) (*domain.Article, error)
    DeleteArticle(idArticle int) error
    GetArticleById(idArticle int) (*domain.Article, error)
    GetAllArticle() ([]*domain.Article, error)
}

func NewArticleService(articleRepository repository.IArticleRepository) IArticleService {
    return &articleService{repo: articleRepository}
}

type articleService struct {
    repo repository.IArticleRepository
}

func (service articleService) CreateArticle(article *dto.Article) (*domain.Article, error) {
    articleToCreate, err := service.repo.Create(article)
    if err != nil {
        log.Error(err)
        return nil, err
    }
    return articleToCreate, nil
}

func (service articleService) DeleteArticle(idArticle int) error {
    err := service.repo.Delete(idArticle)
    if err != nil {
        log.Error(err)
        return err
    }
    return nil
}

func (service articleService) GetArticleById(idArticle int) (*domain.Article, error) {
    article, err := service.repo.FindById(idArticle)
    if err != nil {
        log.Error(err)
        return nil, err
    }
    return article, nil
}

func (service articleService) GetAllArticle() ([]*domain.Article, error) {
    articles, err := service.repo.FindAll()
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return articles, nil
}
