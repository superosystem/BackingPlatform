package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superosystem/BackingPlatform/backend/src/dto"
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/middleware"
	"github.com/superosystem/BackingPlatform/backend/src/model"
	"github.com/superosystem/BackingPlatform/backend/src/service"
)

type transactionController struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(routers *gin.Engine, service service.TransactionService, userService service.UserService, authService middleware.JWT) *transactionController {
	controller := &transactionController{service}

	router := routers.Group("/api/v1/transaction")
	{
		router.GET("/:id/campaign", middleware.AuthMiddleware(authService, userService), controller.GetCampaignTransactions)
		router.GET("", middleware.AuthMiddleware(authService, userService), controller.GetUserTransactions)
		router.POST("", middleware.AuthMiddleware(authService, userService), controller.CreateTransaction)
		router.POST("/notification", controller.GetNotification)
	}

	return controller
}

func (controller *transactionController) GetCampaignTransactions(c *gin.Context) {
	var input model.GetCampaignTransactionsRequest

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := model.RestResult("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	input.User = currentUser

	transactions, err := controller.transactionService.GetTransactionsByCampaignID(input)
	if err != nil {
		response := model.RestResult("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.RestResult("Campaign's transactions", http.StatusOK, "success", dto.CampaignTransactionsDTO(transactions))
	c.JSON(http.StatusOK, response)
}

func (controller *transactionController) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	transactions, err := controller.transactionService.GetTransactionsByUserID(userID)
	if err != nil {
		response := model.RestResult("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.RestResult("User's transactions", http.StatusOK, "success", dto.UserTransactionsDTO(transactions))
	c.JSON(http.StatusOK, response)
}

func (controller *transactionController) CreateTransaction(c *gin.Context) {
	var input model.CreateTransactionRequest

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := model.ErrorValidators(err)
		errorMessage := gin.H{"errors": errors}

		response := model.RestResult("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	input.User = currentUser

	newTransaction, err := controller.transactionService.CreateTransaction(input)

	if err != nil {
		response := model.RestResult("Failed to create transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	response := model.RestResult("Success to create transaction", http.StatusOK, "success", dto.TransactionDTO(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (controller *transactionController) GetNotification(c *gin.Context) {
	var input model.TransactionNotifyRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := model.RestResult("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	err = controller.transactionService.ProcessPayment(input)
	if err != nil {
		response := model.RestResult("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}
	c.JSON(http.StatusOK, input)
}
