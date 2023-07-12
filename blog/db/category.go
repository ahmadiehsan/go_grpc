package db

import (
	"go_grpc/share/db"
)

type Category struct {
	db.BaseModel
	Name string
}
