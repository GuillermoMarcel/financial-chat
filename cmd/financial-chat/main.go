package main

import (
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/repositories"
	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.Default()

	userRepo := repositories.UserRepo{}



	loginController := routers.LoginController{
		Log: logger,
		Repo: &userRepo,
	}

	viewController := routers.ViewsController{}

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", loginController.Login)
	}

	r.LoadHTMLGlob("../../internal/financial-chat/views/*.html")
	r.Static("/assets", "../../internal/financial-chat/assets")
	view := r.Group("/")
	{
		view.GET("/", viewController.Login)
		view.GET("/chatroom", viewController.Home)
	}

	r.Run()
}