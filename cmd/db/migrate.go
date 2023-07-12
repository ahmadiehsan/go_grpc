package main

import (
	blogdb "gogrpc/blog/db"
	sharedb "gogrpc/share/db"
)

var models = []interface{}{
	&blogdb.Article{},
	&blogdb.Category{},
}

func main() {
	db := sharedb.ConnectDB()
	if err := db.AutoMigrate(models...); err != nil {

	}
}
