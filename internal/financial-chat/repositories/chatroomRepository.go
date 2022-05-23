package repositories

import (
	"time"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatroomRepo struct {
	DB *gorm.DB
}

func (r ChatroomRepo) FindChatroom(chatId string) models.Chatroom {
	return models.Chatroom{}
}

func (r ChatroomRepo) SaveMessage(message string, user models.User, chat models.Chatroom) {
	m := &models.Message{
		MessageId: uuid.New().String(),
		Timestamp: time.Now(),
		Content:   message,
		Sender:    user,
		Chatroom: chat,
	}

	r.DB.Create(&m)
}

func (r ChatroomRepo) GetLatestMessages(chatroomId uint) []models.Message{
	var messages []models.Message
	r.DB.
		Where(&models.Message{ChatroomId: chatroomId}).
		Order("timestamp desc").
		Limit(50).
		Preload("Sender").
		Find(&messages)

	return messages
}