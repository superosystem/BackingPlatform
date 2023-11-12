package application

import (
	"github.com/gin-gonic/gin"
	"github.com/superosystem/BackingPlatform/backend/src/controller"
)

func StartApiV1(server *gin.Engine, service *Service) {
	controller.NewUserController(server, service.UserService, service.AuthService)
	controller.NewCampaignController(server, service.CampaignService, service.UserService, service.AuthService)
	controller.NewTransactionHandler(server, service.TransactionService, service.UserService, service.AuthService)
}
