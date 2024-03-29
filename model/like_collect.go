package model

import (
	"time"
)

type Like struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       string
	CreatedAt *time.Time `gorm:"type:timestamp"`
}

type Collection struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       string
	CreatedAt *time.Time `gorm:"type:timestamp"`
}
