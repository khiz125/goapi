package infrastructure

import (
  "gorm.io/gorm"

  "goapi/domain"
)

type GormArticleRepository struct {
  db *gorm.DB
}

func NewGormArticleRepository(db *gorm.DB) *GormArticleRepository {
  return &GormArticleRepository{db: db}
}

func (r *GormArticleRepository) FindByID(id domain.ArticleIdVal) (domain.Article, error) {
  var article domain.Article
  if err := r.db.First(&article, id).Error; err != nil {
    return domain.Article{}, err
  }
  return article, nil
}

func (r *GormArticleRepository) FindAll() ([]domain.Article, error) {
  var article []domain.Article
  if err := r.db.Find(&articles, id).Error; err != nil {
    return domain.Article{}, err
  }
  return articles, nil
}

func (r *GormArticleRepository) Create(article domain.Article) error {
  return r.db.Create(&article).Error
}

func (r *GormArticleRepository) Update(article *domain.Article) error {
  if err := r.db.Model(&domain.Article{}).Where("id = ?", article.ID.Getvalue()).Update(article).Error; err != nil {
    return err
  }
  var updatedArticle domain.Article
  if err := r.db.First(&updatedArticle, article.ID.GetValue()).Error; err != nil {
    return nil, err
  }
  return &updatedArticle, nil
}

func (r *GormArticleRepository) Delete(id domain.ArticleIdVal) error {
  if err := r.db.Delete(&domain.Article{}, id).Error; err != nil {
    return err
  }
  return nil
}
