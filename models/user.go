package models

import (
	"time"
)

type User struct {
	ID           uint        `gorm:"primaryKey"`
	DepartmentID *uint       `gorm:"column:department_id"`
	Name         string      `gorm:"column:name"`
	Email        string      `gorm:"column:email"`
	Password     *string     `gorm:"column:password" json:"-"`
	DeletedAt    *time.Time  `gorm:"column:deleted_at" json:"-"`
	Department   *Department `gorm:"foreignKey:DepartmentID" json:"Department,omitempty"`
}

func (User) TableName() string {
	return "users"
}
