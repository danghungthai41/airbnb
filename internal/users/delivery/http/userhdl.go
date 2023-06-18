package userhttp

import (
	usermodel "airbnb-golang/internal/users/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type IUserHdl interface {
	CreateUser(context.Context, *usermodel.User) (*usermodel.User, error)
}

type UserHdl struct {
	userHdl IUserHdl
}

func NewUserHanlder(hdl IUserHdl) *UserHdl {
	return &UserHdl{hdl}
}

func (hdl *UserHdl) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user usermodel.User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := hdl.userHdl.CreateUser(c.Request.Context(), &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		c.JSON(http.StatusOK, gin.H{"data": data})

	}
}
