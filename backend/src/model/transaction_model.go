package model

import (
	"time"

	"github.com/superosystem/BackingPlatform/backend/src/entity"
)

type GetCampaignTransactionsRequest struct {
	ID   int `uri:"id" binding:"required"`
	User entity.User
}

type CreateTransactionRequest struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       entity.User
}

// For Recipe Notify from Payment Gateway
type TransactionNotifyRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

type TransactionResponse struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

type UserTransactionResponse struct {
	ID        int                  `json:"id"`
	Amount    int                  `json:"amount"`
	Status    string               `json:"status"`
	CreatedAt time.Time            `json:"created_at"`
	Campaign  CampaignUserResponse `json:"campaign"`
}

type CampaignTransactionResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
