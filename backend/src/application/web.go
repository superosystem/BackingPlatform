package application

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/superosystem/BackingPlatform/backend/src/middleware"
	webHandler "github.com/superosystem/BackingPlatform/backend/src/web/handler"
)

func StartWeb(routers *gin.Engine, service *Service) {

	// Load html/tmpl
	routers.HTMLRender = loadTemplates("./src/web/templates")

	// STATIC FILE like css, js, fonts,
	routers.Static("/css", "./src/web/assets/css")
	routers.Static("/js", "./src/web/assets/js")
	routers.Static("/webfonts", "./src/web/assets/webfonts")

	userWebHandler := webHandler.NewUserHandler(service.UserService)
	campaignWebHanlder := webHandler.NewCampaignHandler(service.CampaignService, service.UserService)
	transactionWebHandler := webHandler.NewTransactionHandler(service.TransactionService)
	sessionWebHandler := webHandler.NewSessionHandler(service.UserService)

	router := routers.Group("/")
	{
		router.GET("/users", middleware.AuthAdminMiddleware(), userWebHandler.Index)
		router.GET("/users/new", userWebHandler.New)
		router.POST("/users", userWebHandler.Create)
		router.GET("/users/edit/:id", userWebHandler.Edit)
		router.POST("/users/update/:id", middleware.AuthAdminMiddleware(), userWebHandler.Update)
		router.GET("/users/avatar/:id", middleware.AuthAdminMiddleware(), userWebHandler.NewAvatar)
		router.POST("/users/avatar/:id", middleware.AuthAdminMiddleware(), userWebHandler.CreateAvatar)

		router.GET("/campaigns", middleware.AuthAdminMiddleware(), campaignWebHanlder.Index)
		router.GET("/campaigns/new", middleware.AuthAdminMiddleware(), campaignWebHanlder.New)
		router.POST("/campaigns", middleware.AuthAdminMiddleware(), campaignWebHanlder.Create)
		router.GET("/campaigns/image/:id", middleware.AuthAdminMiddleware(), campaignWebHanlder.NewImage)
		router.POST("/campaigns/image/:id", middleware.AuthAdminMiddleware(), campaignWebHanlder.CreateImage)
		router.GET("/campaigns/edit/:id", middleware.AuthAdminMiddleware(), campaignWebHanlder.Edit)
		router.POST("/campaigns/update/:id", middleware.AuthAdminMiddleware(), campaignWebHanlder.Update)
		router.GET("/campaigns/show/:id", middleware.AuthAdminMiddleware(), campaignWebHanlder.Show)
		router.GET("/transactions", middleware.AuthAdminMiddleware(), transactionWebHandler.Index)

		router.GET("/login", sessionWebHandler.New)
		router.POST("/session", sessionWebHandler.Create)
		router.GET("/logout", sessionWebHandler.Destroy)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
