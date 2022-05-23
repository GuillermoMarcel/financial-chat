package repositories

import (
	"log"
	"time"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatroomRepo struct {
	db *gorm.DB
}

func (r ChatroomRepo) FindChatroom(chatId string) models.Chatroom {
	return models.Chatroom{}
}

func (r ChatroomRepo) SaveMessage(message string, user models.User, chat models.Chatroom) {
	m := &models.Message{
		MessageId: uuid.New().String(),
		Timestamp: time.Now(),
		Content: message,
		Sender: user,
		
	}

	log.Printf("new message %v\n", m)


}
