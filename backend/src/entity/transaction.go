package entity

import "time"

type Transaction struct {
	ID         int `gorm:"primaryKey"`
	UserID     int `gorm:"not null"`
	CampaignID int `gorm:"not null"`
	Amount     int
	Status     string    `gorm:"type:VARCHAR(50)"`
	Code       string    `gorm:"type:VARCHAR(50)"`
	PaymentURL string    `gorm:"type:VARCHAR(255)"`
	User       User      `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	Campaign   Campaign  `gorm:"foreignKey:CampaignID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime"`
}
