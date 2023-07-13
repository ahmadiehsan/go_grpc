package blog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleService struct {
	DB *gorm.DB
}

func (as *ArticleService) Create(a *Article) *Article {
	as.DB.Create(a)
	return a
}

func (as *ArticleService) Update(a *Article) *Article {
	as.DB.Save(a)
	return a
}

func (as *ArticleService) Delete(id uuid.UUID) {
	as.DB.Delete(&Article{}, id)
}

func (as *ArticleService) Get(id uuid.UUID) (*Article, error) {
	var a Article
	result := as.DB.Take(&a)
	if result.Error != nil {
		return nil, result.Error
	}

	return &a, nil
}
