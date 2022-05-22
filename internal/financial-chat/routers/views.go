package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ViewsController struct {

}


func (vc ViewsController) Chatroom(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok {
		vc.Login(c)
		return	
	}


	user, ok := c.GetQuery("us")
	if !ok {
		vc.Login(c)
	}

	c.HTML(
		http.StatusOK,
		"chatroom.html",
		gin.H{
			"chatroomId": id,
			"chatroomName": id, 
			"username": user,
		},
	)
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
		"login.html",
		gin.H{},
	)
}