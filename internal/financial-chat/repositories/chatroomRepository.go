package repositories

import (
	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"gorm.io/gorm"
)

type ChatroomRepo struct {
	db *gorm.DB
}

func (r ChatroomRepo) FindChatroom(chatId string) models.Chatroom {
	return models.Chatroom{}
}

func (r ChatroomRepo) SendMessage() {

}
