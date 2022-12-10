package router

import (
	"fintech/internal/handler"
	"fintech/pkg/errors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"error": "invalid route"})
	})

	router.POST("/", h.GenerateShortURL, errors.Middleware())
	router.GET("/:short", h.GetURLByShort, errors.Middleware())

	return router
}
