package repositories

import (
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	Log *log.Logger
	DB *gorm.DB
}

func (r UserRepo) LoginUser(username string, password string) *models.User {

	var result models.User
	r.DB.Where(&models.User{Username: username}).Preload("Chatrooms").First(&result)

	if result.Username =="" {
		return nil
	}
	if result.Password == password {
		return &result
	}
	return nil
}

func (r UserRepo) FindUser(username string) *models.User {
	return &models.User{
		Username: username,
		Password: "***",
		Name:     "Juan",
	}
}
