package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name   string
	Email  string
	Role   string `gorm:"default:user"`
	Img    string
	Active bool `gorm:"default:false"`
}
