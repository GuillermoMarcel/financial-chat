package models

import "time"

type Chatroom struct {
	Id      string
	Name    string
	Members []User
	Messages []Message
}

type Message struct {
	Timestamp time.Time
	Content string 
	Sender User
}