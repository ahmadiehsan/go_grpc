package blog

import (
	"gogrpc/internal/database"
)

type Article struct {
	database.BaseModel
	Title      string     `json:"title"`
	Content    string     `gorm:"type:text" json:"content"`
	Categories []Category `gorm:"many2many:article_categories" json:"categories"`
}

type Category struct {
	database.BaseModel
	Name string `json:"name"`
}
