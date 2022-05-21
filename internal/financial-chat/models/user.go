package models

type User struct {
	userId    string
	Username  string
	Password  string
	Name      string
	Chatrooms []Chatroom
}