package routers

import (
	"log"
	"net/http"

	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/models"
	"github.com/GuillermoMarcel/financial-chat/internal/financial-chat/repositories"
	"github.com/gin-gonic/gin"
)

const (
	MissingFields = "Missing fields"
	ErrInput      = "Could not read input"
	NotAuthorized = "Not Authorized"
	MessageField  = "message"
	UserField     = "user"
)

type LoginController struct {
	Log  *log.Logger
	Repo *repositories.UserRepo
}

func (lc LoginController) Login(c *gin.Context) {

	var request models.User

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MessageField: ErrInput,
		})
		return
	}

	if request.Username == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			MessageField: MissingFields,
		})
		return
	}

	user := lc.Repo.LoginUser(request.Username, request.Password)

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			MessageField: NotAuthorized,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		UserField: user,
	})
}
