package routes

import (
	"github.com/RianIhsan/shorten_url/feature/url"
	"github.com/gin-gonic/gin"
)

func URlRoute(r *gin.Engine, h url.URLHandlerInterface) {
	urlGroup := r.Group("/v1")
	{
		urlGroup.POST("/short", h.CreateURL)
		urlGroup.GET("/r/:redirect", h.RedirectURL)
	}
}
