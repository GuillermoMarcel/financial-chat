package main

import (
	"errors"
	"log"
	"os"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openDatabase(dbFileName string, init bool) (*gorm.DB, error) {

	if init {
		os.Remove(dbFileName)
	}

	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		return nil, errors.New("unable to connect to database")
	}

	if init {
		err = initializeDatabase(db)
		if err != nil {
			return nil, err
		}
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

	user1 := &models.User{
		Username: "guille",
		Password: "guille",
		Name:     "Guillermo",
	}

	user2 := &models.User{
		Username: "andre",
		Password: "asdf",
		Name:     "Andrea",
	}

	chat1 := &models.Chatroom{
		Name:    "Chatroom1",
		Members: []*models.User{user1, user2},
	}

	chat2 := &models.Chatroom{
		Name:    "Other Chatroom",
		Members: []*models.User{user1, user2},
	}

	db.Create(user1)
	db.Create(user2)

	db.Create(chat1)
	db.Create(chat2)

	return nil
}
