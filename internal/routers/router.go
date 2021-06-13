package routers

import (
	"time"

	"github.com/tianmai777/blog/global"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tianmai777/blog/docs"
	"github.com/tianmai777/blog/internal/middleware"
	"github.com/tianmai777/blog/internal/routers/api"
	v1 "github.com/tianmai777/blog/internal/routers/api/v1"
	"github.com/tianmai777/blog/pkg/limiter"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Translations())
	r.Use(middleware.JWT())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))

	// swag router
	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	r.POST("/auth", api.GetAuth)

	return r
}
