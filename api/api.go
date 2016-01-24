package api

import (
	db "github.com/dancannon/gorethink"
	"github.com/gophergala2016/ring_leader/resources"
	"github.com/gophergala2016/ring_leader/services"
	"github.com/gin-gonic/gin"
)

type API struct {
	DB *db.Session
}

/*
	Policies
*/

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

/*
	Resources
*/

func (a *API) createResource(c *gin.Context) {
	res, err := resources.UnmarshalJSON(c.Request)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	s := services.ResourceService{}
	err = s.InsertResource(a.DB, res)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "")
}

func (a *API) getResources(c *gin.Context) {
	s := services.ResourceService{}
	resType := c.Param("type")
	res, err := s.GetResources(a.DB, resType)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"resources": res})
}

func (a *API) deleteResource(c *gin.Context) {
	c.JSON(200, "")
}

func (a *API) deleteResources(c *gin.Context) {
	c.JSON(200, "")
}

func (a *API) updateResource(c *gin.Context) {
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
		res := v1.Group("resources")
		{
			res.POST("/", api.createResource)
			res.GET("/:type", api.getResources)
			//res.GET("/:type/:id", api.getResources)
			res.DELETE("/", api.deleteResources)
			res.DELETE("/:id", api.deleteResource)
			res.PUT("/:id", api.updateResource)
		}
	}

}
