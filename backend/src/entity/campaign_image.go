package entity

import "time"

type CampaignImage struct {
	ID         int       `gorm:"primaryKey"`
	CampaignID int       `gorm:"not null"`
	FileName   string    `gorm:"type:VARCHAR(255)"`
	IsPrimary  int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime"`
}
