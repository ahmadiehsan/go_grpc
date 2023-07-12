package blogdb

import sharedb "gogrpc/share/db"

type Category struct {
	sharedb.BaseModel
	Name string
}
