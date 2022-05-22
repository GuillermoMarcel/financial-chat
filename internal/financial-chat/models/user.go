package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Username  string
	Password  string
	Name      string
	Chatrooms []*Chatroom `gorm:"many2many:user_chatrooms;"`
}
