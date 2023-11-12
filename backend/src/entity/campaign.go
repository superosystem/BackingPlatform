package entity

import "time"

type Campaign struct {
	ID               int    `gorm:"primaryKey"`
	UserID           int    `gorm:"not null"`
	Name             string `gorm:"type:VARCHAR(100);not null"`
	ShortDescription string `gorm:"type:VARCHAR(100)"`
	Description      string `gorm:"type:TEXT"`
	Perks            string `gorm:"type:TEXT"`
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string    `gorm:"type:TEXT"`
	CreatedAt        time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"not null;autoUpdateTime"`
	CampaignImages   []CampaignImage
	User             User `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
}
