package repository

import (
    "time"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/model/domain"
    "github.com/itp-backend/backend-a-co-create/model/dto"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type IArticleRepository interface {
    Create(article *dto.Article) (*domain.Article, error)
    Delete(idArticle int) error
    FindById(idArticle int) (*domain.Article, error)
    FindAll() ([]*domain.Article, error)
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
    return articleRepository{DB: db}
}

type articleRepository struct {
    DB *gorm.DB
}

func (repo articleRepository) Create(article *dto.Article) (*domain.Article, error) {
    a := &domain.Article{
        PostingDate: time.Now().UnixMilli(),
        Kategori:    article.Kategori,
        Judul:       article.Judul,
        IsiArtikel:  article.IsiArtikel,
        IdUser:      article.IdUser,
    }

    result := repo.DB.Create(&a)
    if result.Error != nil {
        log.Error(result.Error)
        return nil, result.Error
    }
    return a, nil
}

func (repo articleRepository) Delete(idArticle int) error {
    var article domain.Article
    article.IdArtikel = idArticle

    err := repo.DB.Where("id_artikel = ?", idArticle).First(&article).Error

    switch err {
    case nil:
        repo.DB.Delete(&article)
        return nil
    case gorm.ErrRecordNotFound:
        return errors.NewInternalError(err, "Error: not found")
    default:
        return errors.NewInternalError(err, "Error: database error")
    }
}

func (repo articleRepository) FindById(idArticle int) (*domain.Article, error) {
    var article domain.Article
    article.IdArtikel = idArticle

    if err := repo.DB.First(&article).Error; err != nil {
        log.Error(err)
        return &article, err
    }

    return &article, nil
}

func (repo articleRepository) FindAll() ([]*domain.Article, error) {
    var articles []*domain.Article
    if err := repo.DB.Table("articles").Find(&articles).Error; err != nil {
        log.Error(err)
        return articles, err
    }

    return articles, nil
}


