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
	Service wschat.ChatroomService
	Hub *wschat.Hub
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
	
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &wschat.Client{
		Hub: r.Hub, 
		Conn: conn, 
		Send: make(chan []byte, 256),
	}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}

func (r ChatRoomRouter) SendMessage(c *gin.Context){
	
}
