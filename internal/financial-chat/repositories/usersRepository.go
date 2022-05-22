package repositories

import (
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
)

type UserRepo struct {
	log *log.Logger
}

func (r UserRepo) LoginUser(username string, password string) *models.User {
	if username == "nil" {
		return nil
	}
	return &models.User{
		Username: username,
		Password: "***",
		Name:     "Juan",
	}
}

func (r UserRepo) FindUser(username string) *models.User {
	return &models.User{
		Username: username,
		Password: "***",
		Name:     "Juan",
	}
}
