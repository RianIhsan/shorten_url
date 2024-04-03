package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RianIhsan/shorten_url/feature/url"
	"github.com/RianIhsan/shorten_url/feature/url/dto"
	"github.com/gin-gonic/gin"
)

type urlHandler struct {
	urlSvc url.URLServiceInterface
}

func NewURLHandler(urlSvc url.URLServiceInterface) url.URLHandlerInterface {
	return &urlHandler{
		urlSvc: urlSvc,
	}
}

func (h *urlHandler) CreateURL(c *gin.Context) {
	var req dto.CreateURLRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid payload",
		})
		return
	}

	url, err := h.urlSvc.CreateURL(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed create shortlink"+err.Error(),
		})
		return
	}
	baseURL := "http://127.0.0.1:8808/v1/r/"
	shortURLString := fmt.Sprintf("%s%s", baseURL, url.ShorterURL)

	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURLString,
	})
}
func (h *urlHandler) RedirectURL(c *gin.Context) {
	param := c.Param("redirect")
	url, err := h.urlSvc.GetShortURL(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed get shortlink",
		})
		return
	}
	if !strings.HasPrefix(url.OriginalURL, "http://") && !strings.HasPrefix(url.OriginalURL, "https://") {
		url.OriginalURL = "http://" + url.OriginalURL
	}
	c.Redirect(http.StatusTemporaryRedirect, url.OriginalURL)
}
