package blog

import "gogrpc/pkg/db"

type Article struct {
	db.BaseModel
	Title      string
	Content    string     `gorm:"type:text"`
	Categories []Category `gorm:"many2many:article_categories;"`
}

type Category struct {
	db.BaseModel
	Name string
}
