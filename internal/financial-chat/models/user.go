package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string `json:"-"`
	Name      string
	Chatrooms []*Chatroom `gorm:"many2many:user_chatrooms;"`
}