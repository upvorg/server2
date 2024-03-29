package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	ParentID  uint
	TargetID  uint
	Uid       uint           `gorm:"not null"`
	Pid       string         `gorm:"not null"`
	Vid       uint           `gorm:"default:1"`
	Content   string         `gorm:"size:200"`
	Color     string         `gorm:"size:10"`
	CreatedAt *time.Time     `gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Creator   *User          `gorm:"foreignKey:Uid" json:"Creator,omitempty"`
	Post      *Post          `gorm:"foreignKey:Pid" json:"Post,omitempty"`
	Children  []*Comment     `gorm:"foreignKey:parent_id" json:"Children,omitempty"`
}
