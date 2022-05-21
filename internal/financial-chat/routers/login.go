package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Log *log.Logger
}

func (lc LoginController) Login(c *gin.Context){
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Not implemented",
	})
}