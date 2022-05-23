package routers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/repositories"
	"github.com/gin-gonic/gin"
)

type ViewsController struct {
	ChatroomRepo repositories.ChatroomRepo
}

func (vc ViewsController) Chatroom(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		vc.Login(c)
		return
	}

	user, ok := c.GetQuery("us")
	if !ok {
		vc.Login(c)
	}

	chatId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad format"})
	}

	chat := vc.ChatroomRepo.FindChatroom(uint(chatId))

	c.HTML(
		http.StatusOK,
		"chatroom.html",
		gin.H{
			"chatroomId":   chat.ID,
			"chatroomName": chat.Name,
			"username":     user,
		},
	)
}

func (vc ViewsController) Login(c *gin.Context) {

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
