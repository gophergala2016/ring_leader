package main

import (
	"fmt"
	db "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/gophergala2016/ring_leader/api"
	"github.com/satori/go.uuid"
	"log"
	"os"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(RequestIdMiddleware())
	DB, err := db.Connect(db.ConnectOpts{
		Address: fmt.Sprintf("%s:%d", os.Getenv("RETHINKDB_PORT_28015_TCP_ADDR"), 28015),
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	api.Init(r, DB)
	r.GET("/ping", func(c *gin.Context) {
		if DB.IsConnected() {
			c.String(200, "ok")
		} else {
			c.String(500, "not ok")
		}
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
