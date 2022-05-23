package wschat

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/repositories"
)

type ChatroomService struct {
	log          *log.Logger
	chatroomRepo *repositories.ChatroomRepo
	userRepo     *repositories.UserRepo
	hubs         map[uint]*hub
	incoming     chan incomingMessage
}

func NewChatroomService(log *log.Logger, chatRepo *repositories.ChatroomRepo, userRepo *repositories.UserRepo) *ChatroomService {
	log.SetPrefix("ChatService")
	hubs := make(map[uint]*hub)
	incChan := make(chan incomingMessage)
	cs := ChatroomService{
		hubs:         hubs,
		chatroomRepo: chatRepo,
		userRepo:     userRepo,
		log:          log,
		incoming:     incChan,
	}
	go cs.readIncoming()
	return &cs
}

func (s ChatroomService) initChatroom(chatId uint) *hub {
	s.log.Printf("initializing chatroom, chatId: %d\n", chatId)
	hub := newHub()
	go hub.run()
	s.hubs[chatId] = hub
	return hub
}

func (s ChatroomService) RegisterIncoming(w http.ResponseWriter, r *http.Request, chatroomId uint, userId string) error {

	user := s.userRepo.FindUser(userId)
	var chatroom models.Chatroom
	member := false
	for _, c := range user.Chatrooms {
		if c.ID == chatroomId {
			member = true
			chatroom = *c
			break
		}
	}
	if !member {
		return errors.New("not a member of chatroom")
	}

	hub, ok := s.hubs[chatroomId]

	if !ok {
		hub = s.initChatroom(chatroomId)
	}

	if hub == nil {
		return errors.New("not able to start chatroom")
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	client := &Client{
		Hub:    hub,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		recive: s.incoming,
		user:   user,
		chatroom: &chatroom,
	}
	client.Hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()

	s.loadOldMessages(client, chatroom)

	return nil
}

func (s ChatroomService) loadOldMessages(client *Client, chatroom models.Chatroom){
	messages := s.chatroomRepo.GetLatestMessages(chatroom.ID)

	for i := len(messages)-1; i >= 0; i-- {
		m := messages[i]
		c := fmt.Sprintf("(%s) %s: %s", m.Timestamp.Format("2006-01-02 15:04:05"), m.Sender.Name, m.Content)
		client.Send <- []byte(c)
	 }
}


func (s ChatroomService) readIncoming() {
	defer close(s.incoming)
	for {
		message := <-s.incoming
		content := string(message.content)

		s.chatroomRepo.SaveMessage(content, *message.user, *message.chatroom)

		outgoing := fmt.Sprintf("(%s) %s: %s", time.Now().Format("2006-01-02 15:04:05"), message.user.Name, content)

		message.client.Hub.broadcast <- []byte(outgoing)
	}
}

