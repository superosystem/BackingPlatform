package application

import (
	"github.com/superosystem/BackingPlatform/backend/src/config"
	"github.com/superosystem/BackingPlatform/backend/src/repository"
)

type Repository struct {
	UserRepository        repository.UserRepository
	CampaignRepository    repository.CampaignRepository
	TransactionRepository repository.TransactionRepository
}

func NewRepository(conn *config.Connect) *Repository {
	return &Repository{
		UserRepository:        repository.NewUserRepository(conn.MySQL),
		CampaignRepository:    repository.NewCampaignRepository(conn.MySQL),
		TransactionRepository: repository.NewTransactionRepository(conn.MySQL),
	}
}
