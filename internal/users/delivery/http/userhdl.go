package userhttp

import (
	usermodel "airbnb-golang/internal/users/model"
	"airbnb-golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type IUserUC interface {
	Register(context.Context, *usermodel.UserCreate) error
	Login(context.Context, *usermodel.UserLogin) (*utils.Token, error)
}

type UserHdl struct {
	userUC IUserUC
}

func NewUserHanlder(uc IUserUC) *UserHdl {
	return &UserHdl{uc}
}

func (hdl *UserHdl) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := hdl.userUC.Register(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		c.JSON(http.StatusOK, gin.H{"data": data.Id})

	}
}
func (hdl *UserHdl) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials usermodel.UserLogin

		if err := c.ShouldBind(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := hdl.userUC.Login(c.Request.Context(), &credentials)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": token})

	}
}
