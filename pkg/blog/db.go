package blog

import (
	"gogrpc/internal/database"
)

type Article struct {
	database.BaseModel
	Title      string
	Content    string     `gorm:"type:text"`
	Categories []Category `gorm:"many2many:article_categories;"`
}

type Category struct {
	database.BaseModel
	Name string
}
