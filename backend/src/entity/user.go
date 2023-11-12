package entity

import "time"

type User struct {
	ID             int       `gorm:"primaryKey"`
	Name           string    `gorm:"type:VARCHAR(100);not null"`
	Occupation     string    `gorm:"type:VARCHAR(100)"`
	Email          string    `gorm:"type:VARCHAR(100);unique;not null;"`
	Password       string    `gorm:"not null"`
	AvatarFileName string    `gorm:"type:VARCHAR(255)"`
	Role           string    `gorm:"type:VARCHAR(100)"`
	Token          string    `gorm:"type:VARCHAR(255)"`
	CreatedAt      time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"not null;autoUpdateTime"`
}
