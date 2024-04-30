package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string
	Email     string
	Password  string
	Role      string `gorm:"default:user"`
	Img       string
	Active    bool      `gorm:"default:false"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
