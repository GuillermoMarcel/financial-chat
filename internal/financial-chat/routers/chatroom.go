package routers

import (
	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/repositories"
	"github.com/gin-gonic/gin"
)

type ChatRoomRouter struct {
	repo repositories.ChatroomRepo
}

func (r ChatRoomRouter) OpenChatroom(c *gin.Context){

}

func (r ChatRoomRouter) SendMessage(c *gin.Context){
	
}