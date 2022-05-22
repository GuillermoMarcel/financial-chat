package wschat

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/repositories"
)

type ChatroomService struct {
	log          *log.Logger
	chatroomRepo *repositories.ChatroomRepo
	userRepo     *repositories.UserRepo
	hubs         map[string]*hub
	incoming     chan incomingMessage
}

func NewChatroomService(log *log.Logger, chatRepo *repositories.ChatroomRepo, userRepo *repositories.UserRepo) *ChatroomService {
	log.SetPrefix("ChatService")
	hubs := make(map[string]*hub)
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

func (s ChatroomService) initChatroom(chatId string) *hub {
	s.log.Printf("initializing chatroom, chatId: %s\n", chatId)
	hub := newHub()
	go hub.run()
	s.hubs[chatId] = hub
	return hub
}

func (s ChatroomService) RegisterIncoming(w http.ResponseWriter, r *http.Request, chatroomId string, userId string) error {

	user := s.userRepo.FindUser(userId)

	member := false
	for _, c := range user.Chatrooms {
		if c.ChatroomId == chatroomId {
			member = true
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
	}
	client.Hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()

	return nil
}

func (s ChatroomService) readIncoming() {
	defer close(s.incoming)
	for {
		message := <-s.incoming
		content := string(message.content)

		outgoing := fmt.Sprintf("%s: %s", message.user.Name, content)

		message.client.Hub.broadcast <- []byte(outgoing)
	}
}
