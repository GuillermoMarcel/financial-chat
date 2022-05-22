package main

import (
	"errors"
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../../database.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("unable to connect to database")
	}

	err = initializeDatabase(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initializeDatabase(db *gorm.DB) error {
	// Migrate the schema
	err := db.AutoMigrate(&models.User{}, &models.User{}, &models.Message{})
	if err != nil {
		log.Println(err)
		return err
	}

	user1 :=&models.User{
		ID: "us-01",
		Username: "guille",
		Password: "guille",
		Name:     "Guillermo",
	}

	user2 := &models.User{
		ID: "us-02",
		Username: "andre",
		Password: "asdf",
		Name:     "Andrea",
	}

	chat1 := &models.Chatroom{
		ChatroomId: "cr-001",
		Name:       "The very first",
		Members: []*models.User{user1, user2},
	}

	chat2 := &models.Chatroom{
		ChatroomId: "cr-002",
		Name:       "Second best",
		Members: []*models.User{user1,user2},
	}

	db.Create(chat1)
	db.Create(chat2)

	db.Create(user1)
	db.Create(user2)

	return nil

}
