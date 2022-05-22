package models

import (
	"time"

	"gorm.io/gorm"
)

type Chatroom struct {
	gorm.Model
	Name       string
	Members    []*User    `gorm:"many2many:user_chatrooms;" json:",omitempty"`
	Messages   []*Message `gorm:"foreignKey:ChatroomId;" json:",omitempty"`
}

type Message struct {
	gorm.Model
	MessageId  string `gorm:"primaryKey"`
	Timestamp  time.Time
	Content    string
	SenderId   string
	Sender     User `gorm:"foreignKey:SenderId"`
	ChatroomId string
}
