package application

import (
	"github.com/superosystem/BackingPlatform/backend/src/middleware"
	"github.com/superosystem/BackingPlatform/backend/src/service"
)

type Service struct {
	UserService        service.UserService
	CampaignService    service.CampaignService
	AuthService        middleware.JWT
	TransactionService service.TransactionService
}

func NewService(r *Repository) *Service {
	return &Service{
		UserService:        service.NewUserService(r.UserRepository),
		CampaignService:    service.NewCampaignService(r.CampaignRepository),
		AuthService:        middleware.NewJwtService(),
		TransactionService: service.NewTransactionService(r.TransactionRepository, r.CampaignRepository, service.NewPaymentService()),
	}
}
