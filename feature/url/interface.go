package url

import (
	"github.com/RianIhsan/shorten_url/entities"
	"github.com/RianIhsan/shorten_url/feature/url/dto"
	"github.com/gin-gonic/gin"
)

type URLRepositoryInterface interface {
	CreateURL(req *entities.MstURL) (*entities.MstURL, error)
	GetShortURL(url string) (*entities.MstURL, error)
}

type URLServiceInterface interface {
	CreateURL(req *dto.CreateURLRequest) (*entities.MstURL, error)
	GetShortURL(url string) (*entities.MstURL, error)
}

type URLHandlerInterface interface {
	CreateURL(c *gin.Context)
	RedirectURL(c *gin.Context)
}
