package blogdb

import sharedb "gogrpc/share/db"

type Article struct {
	sharedb.BaseModel
	Title      string
	Content    string     `gorm:"type:text"`
	Categories []Category `gorm:"many2many:article_categories;"`
}
