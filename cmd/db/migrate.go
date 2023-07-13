package main

import (
	log "github.com/sirupsen/logrus"
	"gogrpc/pkg/blog"
	"gogrpc/pkg/db"
)

var models = []interface{}{
	&blog.Article{},
	&blog.Category{},
}

func main() {
	db := db.ConnectDB()
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal("Error in migration: ", err)
	}
	log.Infof("Done")
}
