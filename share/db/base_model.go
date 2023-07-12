package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TODO
/*func (b *BaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}*/
