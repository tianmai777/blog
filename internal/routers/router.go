package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		// 标签接口
		apiv1.POST("/tags")
		apiv1.GET("/tags")
		apiv1.GET("/tags/:id")
		apiv1.PUT("/tags/:id")
		apiv1.DELETE("/tags/:id")

		// 文章接口
		apiv1.POST("/articles")
		apiv1.GET("/articles")
		apiv1.GET("/articles/:id")
		apiv1.PUT("/articles/:id")
		apiv1.DELETE("/articles/:id")
	}

	return r
}
