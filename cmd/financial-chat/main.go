package main

import (
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/repositories"
	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/routers"
	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/services/wschat"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.Default()

	db, err := initializeDatabase()
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.Printf("databse initialized %s\n", db.Name())

	userRepo := repositories.UserRepo{}
	chatRepo := repositories.ChatroomRepo{}

	chatService := wschat.NewChatroomService(log.Default(), &chatRepo, &userRepo)


	loginController := routers.LoginController{
		Log: logger,
		Repo: &userRepo,
	}
	chatroomController := routers.ChatRoomRouter{
		Logger: logger,
		Service: chatService,
	}

	viewController := routers.ViewsController{}

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", loginController.Login)

		api.Any("/chatroom/ws", chatroomController.OpenChatroom)
	}



	r.LoadHTMLGlob("../../internal/financial-chat/views/*.html")
	r.Static("/assets", "../../internal/financial-chat/assets")
	view := r.Group("/")
	{
		view.GET("/chatroom", viewController.Chatroom)
		view.GET("/", viewController.Login)
	}

	r.Run()
}