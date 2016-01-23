package login

import (
	db "github.com/dancannon/gorethink"
	"github.com/gophergala2016/ring_leader/models"
	"github.com/gophergala2016/ring_leader/services"
	"github.com/gin-gonic/gin"
)

type Login struct {
	DB *db.Session
}
func Init(router *gin.Engine, DB *db.Session) {
	// Simple group: login
	l := &Login{DB}
	loginRouter := router.Group("/login")
	{
		loginRouter.POST("/register", l.registerUser)
	}
}
func (l *Login) registerUser(c *gin.Context) {
	var json models.CreateUser
	if err := c.BindJSON(&json); err != nil {
		c.String(500, err.Error())
		return
	}
	service := &services.UserService{}
	exists := service.UserExistByEmail(l.DB, json.Email)
	if exists == true {
		c.String(500, "user already exists")
		return
	}

	err := service.InsertUser(l.DB, json)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "worked")
}

func (l Login) LoginUser(form models.LoginUser) error {
	return nil
}

func (l Login) ChangeUser(form models.ChangeUser, id int32) error {
	return nil
}
