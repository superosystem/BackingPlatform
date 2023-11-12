package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/superosystem/BackingPlatform/backend/src/dto"
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/middleware"
	"github.com/superosystem/BackingPlatform/backend/src/model"
	"github.com/superosystem/BackingPlatform/backend/src/service"
)

type campaignController struct {
	campaignService service.CampaignService
	userService     service.UserService
	authService     middleware.JWT
}

func NewCampaignController(routers *gin.Engine, campaignService service.CampaignService, userService service.UserService, authService middleware.JWT) *campaignController {
	controller := &campaignController{campaignService, userService, authService}

	router := routers.Group("/api/v1/campaign")
	{
		router.GET("", controller.GetCampaigns)
		router.GET("/:id", controller.GetCampaign)
		router.POST("", middleware.AuthMiddleware(authService, userService), controller.CreateCampaign)
		router.PUT("/:id", middleware.AuthMiddleware(authService, userService), controller.UpdateCampaign)
		router.POST("/image", middleware.AuthMiddleware(authService, userService), controller.UploadImage)
	}

	return controller
}

func (controller *campaignController) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := controller.campaignService.GetCampaigns(userID)
	if err != nil {
		response := model.RestResult("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.RestResult("List of campaigns", http.StatusOK, "success", dto.CampaignsDTO(campaigns))
	c.JSON(http.StatusOK, response)
}

func (controller *campaignController) GetCampaign(c *gin.Context) {
	var input model.GetCampaignDetailRequest

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := model.RestResult("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := controller.campaignService.GetCampaignByID(input)
	if err != nil {
		response := model.RestResult("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if campaignDetail.ID == 0 {
		response := model.RestResult("Campaign is not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := model.RestResult("Campaign detail", http.StatusOK, "success", dto.CampaignDetailDTO(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (controller *campaignController) CreateCampaign(c *gin.Context) {
	var input model.CreateCampaignRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := model.ErrorValidators(err)
		errorMessage := gin.H{"errors": errors}

		response := model.RestResult("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	input.User = currentUser

	newCampaign, err := controller.campaignService.CreateCampaign(input)
	if err != nil {
		response := model.RestResult("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.RestResult("Success to create campaign", http.StatusOK, "success", dto.CampaignDTO(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (controller *campaignController) UpdateCampaign(c *gin.Context) {
	var inputID model.GetCampaignDetailRequest

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := model.RestResult("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData model.CreateCampaignRequest

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := model.ErrorValidators(err)
		errorMessage := gin.H{"errors": errors}

		response := model.RestResult("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)
	inputData.User = currentUser

	updatedCampaign, err := controller.campaignService.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := model.RestResult("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.RestResult("Success to update campaign", http.StatusOK, "success", dto.CampaignDTO(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (controller *campaignController) UploadImage(c *gin.Context) {
	var input model.CreateCampaignImageRequest

	err := c.ShouldBind(&input)
	if err != nil {
		errors := model.ErrorValidators(err)
		errorMessage := gin.H{"errors": errors}

		response := model.RestResult("Failed to upload campaign image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := model.RestResult("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("public/storage/campaigns/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := model.RestResult("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = controller.campaignService.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := model.RestResult("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := model.RestResult("Campaign image successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
