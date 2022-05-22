package routers

import (
	"log"
	"net/http"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/services/wschat"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatRoomRouter struct {
	*log.Logger
	Service *wschat.ChatroomService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (r ChatRoomRouter) OpenChatroom(c *gin.Context){
	chatid, ok := c.GetQuery("chatroom")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "chatroom id requestd"})
		return
	}
	userId, ok := c.GetQuery("user")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user id requestd"})
		return
	}
	r.Logger.Printf("Incoming connection, chat: %s, user: %s", chatid, userId)
	

	err := r.Service.RegisterIncoming(c.Writer, c.Request, chatid, userId)
	if err != nil {
		r.Logger.Printf("error registering: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	
}

func (r ChatRoomRouter) SendMessage(c *gin.Context){
	
}
