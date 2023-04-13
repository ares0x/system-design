package router

import (
	"github.com/gin-gonic/gin"
	"short-url/internal/router/handlers"
	"short-url/internal/router/middlewares"
)

func LoadEngine() *gin.Engine {
	g := gin.Default()
	loadRouter(g)
	loadMiddlewares(g)
	return g
}

func loadMiddlewares(g *gin.Engine) {
	g.Use(middlewares.RequestId())
}

func loadRouter(g *gin.Engine) {
	handler := handlers.NewUrlHandler()
	apiV1Group := g.Group("/api/v1")
	{
		apiV1Group.POST("/api/original", handler.SetOriginalUrl)
		apiV1Group.GET("/api/:shorturl", handler.SetOriginalUrl)
	}
}
