package db

import (
	"go_grpc/share/db"
)

type Article struct {
	db.BaseModel
	Title      string
	Content    string     `gorm:"type:text"`
	Categories []Category `gorm:"many2many:article_categories;"`
}
