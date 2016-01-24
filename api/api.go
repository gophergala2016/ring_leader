package api

import (
	db "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
)

type API struct {
	DB *db.Session
}

func (a *API) createPolicy(c *gin.Context) {
	c.JSON(200, "")
}

func (a *API) getPolicies(c *gin.Context) {
	c.String(200, "policies")
}

func (a *API) deletePolicy(c *gin.Context) {
	c.JSON(200, "")
}

func (a *API) deletePolicies(c *gin.Context) {
	c.JSON(200, "")
}

func (a *API) updatePolicy(c *gin.Context) {
	c.JSON(200, "")
}

func Init(router *gin.Engine, DB *db.Session, fn func() gin.HandlerFunc) {
	// Simple group: v1
	api := &API{DB}
	v1 := router.Group("/v1")
	{
		v1.Use(fn())
		policies := v1.Group("policies")
		{
			policies.POST("/", api.createPolicy)
			policies.GET("/", api.getPolicies)
			policies.DELETE("/", api.deletePolicies)
			policies.DELETE("/:id", api.deletePolicy)
			policies.PUT("/:id", api.updatePolicy)
		}
	}

}

