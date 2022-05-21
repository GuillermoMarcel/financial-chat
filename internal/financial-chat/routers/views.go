package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ViewsController struct {

}


func (vc ViewsController) Home(c *gin.Context){

}

func (vc ViewsController) Login(c *gin.Context){

	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "time": time.Now()})
        }
    }()

	c.HTML(
		http.StatusOK,
		// "internal/finalcial-chat/views/login.html",
		"login.html",
		gin.H{
			"title": "Web",
			"url":   "./web.json",
		},
	)
}