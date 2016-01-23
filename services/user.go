package services

import (
	db "github.com/dancannon/gorethink"
	"github.com/gophergala2016/ring_leader/models"
	"gopkg.in/hlandau/passlib.v1"
	"log"
)
type UserService struct {
}

func (s *UserService) UserExistByEmail(DB *db.Session, email string) bool {
	res, err := db.Table("users").Filter(db.Row.Field("email").Eq(email)).Count().Run(DB)
	if err != nil {
		return false
	}
	var cnt int
	err = res.One(&cnt)
	defer res.Close()
	if err != nil {
		return false
	}
	return cnt != 0
}

func (s UserService) InsertUser(DB *db.Session, createUser models.CreateUser) error {
	hash, err := passlib.Hash(createUser.Password)
	if err != nil {
		// couldn't hash password for some reason
		return err
	}
	user := models.User{
		Email: createUser.Email,
		Name: createUser.Name,
		Salt: hash,
	}
	_, err = db.Table("users").Insert(user).RunWrite(DB)
	if err != nil {
		return err
	}
	return nil
}
