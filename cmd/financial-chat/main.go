package main

import (
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.Default()

	loginController := routers.LoginController{
		Log: logger,
	}

	r := gin.Default()
	r.GET("/login", loginController.Login)

	r.Run()
}