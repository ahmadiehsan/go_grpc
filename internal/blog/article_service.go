package blog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleService struct {
	DB *gorm.DB
}

func (as *ArticleService) Create(a *Article) error {
	return as.DB.Create(a).Error
}

func (as *ArticleService) Update(a *Article) error {
	return as.DB.Save(a).Error
}

func (as *ArticleService) Delete(id uuid.UUID) error {
	return as.DB.Delete(&Article{}, id).Error
}

func (as *ArticleService) Get(id uuid.UUID) (a *Article, err error) {
	a.ID = id
	err = as.DB.Take(a).Error // TODO use zero log for the recordNotFound error

	return
}
