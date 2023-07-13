package blog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleService struct {
	DB *gorm.DB
}

func (as *ArticleService) Create(a *Article) error {
	result := as.DB.Create(a)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (as *ArticleService) Update(a *Article) error {
	result := as.DB.Save(a)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (as *ArticleService) Delete(id uuid.UUID) error {
	result := as.DB.Delete(&Article{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (as *ArticleService) Get(id uuid.UUID) (*Article, error) {
	var a Article
	a.ID = id
	result := as.DB.Take(&a) // TODO use zero log for the recordNotFound error
	if result.Error != nil {
		return nil, result.Error
	}

	return &a, nil
}
