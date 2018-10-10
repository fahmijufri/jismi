package domain

import (
	"time"
)

type User struct {
	ID         string     `gorm:"not null;column:id;primary_key" json:"id" schema:"id" valid:"uuid,required"`
	FullNumber string     `gorm:"column:full_number" json:"full_number" schema:"full_number" valid:"required"`
	Status     string     `gorm:"not null;column:status" json:"status" schema:"status" valid:"required"`
	CreatedAt  time.Time  `gorm:"not null;column:created_at" json:"created_at" schema:"-"  valid:"optional"`
	UpdatedAt  time.Time  `gorm:"not null;column:updated_at" json:"updated_at" schema:"-"  valid:"optional"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty" schema:"-"  valid:"-"`
}
