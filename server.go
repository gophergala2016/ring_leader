package main

import (
	"fmt"
	db "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gophergala2016/ring_leader/api"
	"github.com/gophergala2016/ring_leader/models"
	"github.com/gophergala2016/ring_leader/login"
	//"github.com/gophergala2016/ring_leader/settings"
	"github.com/satori/go.uuid"
	"encoding/gob"
	"log"
	"os"
	"net/http"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func RequestAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		u := session.Get("user")
		if u == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/ping")
			// TODO Redirect
		}
		log.Printf("%T %v\n", u, u)
		// Validate credentials
		switch user := u.(type) {
		case models.User:
			_ = user
			c.Next()
		default:
			c.Redirect(http.StatusTemporaryRedirect, "/ping")
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(RequestIdMiddleware())
	DB, err := db.Connect(db.ConnectOpts{
		Address: fmt.Sprintf("%s:%d", os.Getenv("RETHINKDB_PORT_28015_TCP_ADDR"), 28015),
		Database: "test",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	db.TableCreate("users").RunWrite(DB)
	store, err := sessions.NewRedisStore(10, "tcp", fmt.Sprintf("%s:%d", os.Getenv("REDIS_1_PORT_6379_TCP_ADDR"), 6379), "", []byte(os.Getenv("REDIS_SECRET")))
	if err != nil {
		log.Fatalln(err.Error())
	}
	gob.Register(models.User{})
	r.Use(sessions.Sessions("session", store))
	api.Init(r, DB, RequestAuthMiddleware)
	login.Init(r, DB)
	r.GET("/ping", func(c *gin.Context) {
		if DB.IsConnected() {
			c.String(200, "ok")
		} else {
			c.String(500, "not ok")
		}
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
