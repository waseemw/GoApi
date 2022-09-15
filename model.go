package main

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        string `gorm:"primaryKey,type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
