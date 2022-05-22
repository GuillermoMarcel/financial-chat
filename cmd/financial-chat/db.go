package main

import (
	"errors"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
	  return nil, errors.New("unable to connect to database")
	}

	  // Migrate the schema
	  db.AutoMigrate(&models.User{})
	  db.AutoMigrate(&models.User{})
	  db.AutoMigrate(&models.Message{})

	return db, nil
}