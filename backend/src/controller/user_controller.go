package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superosystem/BackingPlatform/backend/src/dto"
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/middleware"
	"github.com/superosystem/BackingPlatform/backend/src/model"
	"github.com/superosystem/BackingPlatform/backend/src/service"
)

type userController struct {
	userService service.UserService
	authService middleware.JWT
}

func NewUserController(routers *gin.Engine, userService service.UserService, authService middleware.JWT) *userController {
	controller := &userController{userService, authService}

	router := routers.Group("/api/v1/user")
	{
		router.POST("/register", controller.RegisterUser)
		router.POST("/login", controller.LoginUser)
		router.POST("/email_checkers", controller.CheckEmailAvailability)
		router.POST("/avatar", middleware.AuthMiddleware(authService, userService), controller.UploadAvatar)
		router.GET("/fetch", middleware.AuthMiddleware(authService, userService), controller.FetchUser)
	}

	return controller
}

func (controller *userController) RegisterUser(c *gin.Context) {
	var input model.RegisterRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := model.ErrorValidators(err)
		errorMessage := gin.H{"errors": errors}

		response := model.RestResult("register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := controller.userService.RegisterUser(input)

	if err != nil {
		response := model.RestResult("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := controller.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := model.RestResult("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := dto.UserDTO(newUser, token)
	response := model.RestResult("Account has been registered", http.StatusOK, "success", data)
	c.JSON(http.StatusCreated, response)
}

func (controller *userController) LoginUser(c *gin.Context) {
	var input model.LoginRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := model.ErrorValidators(err)
		errorMessage := gin.H{"errors": errors}

		response := model.RestResult("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := controller.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := model.RestResult("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := controller.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := model.RestResult("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := dto.UserDTO(loggedinUser, token)
	response := model.RestResult("Successfuly loggedin", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (controller *userController) CheckEmailAvailability(c *gin.Context) {
	var input model.CheckEmailRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := model.ErrorValidators(err)
		errorMessage := gin.H{"errors": errors}

		response := model.RestResult("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := controller.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := model.RestResult("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	if !isEmailAvailable {
		response := model.RestResult("Email has been used", http.StatusConflict, "error", data)
		c.JSON(http.StatusConflict, response)
		return
	}

	response := model.RestResult("Email is available", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (controller *userController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := model.RestResult("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	// path file location would be saved
	path := fmt.Sprintf("public/storage/avatars/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := model.RestResult("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = controller.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := model.RestResult("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := model.RestResult("Avatar successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userController) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)

	data := dto.UserDTO(currentUser, "")
	response := model.RestResult("Successfuly fetch user data", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
