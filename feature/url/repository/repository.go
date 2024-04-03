package repository

import (
	"github.com/RianIhsan/shorten_url/entities"
	"github.com/RianIhsan/shorten_url/feature/url"
	"gorm.io/gorm"
)

type urlRepo struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) url.URLRepositoryInterface {
	return &urlRepo{db: db}
}

func (r *urlRepo) CreateURL(req *entities.MstURL) (*entities.MstURL, error) {
	if err := r.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *urlRepo) GetShortURL(url string) (*entities.MstURL, error) {
	var schema entities.MstURL
	res := r.db.Where("shorter_url = ?", url).Find(&schema)
	if res.Error != nil {
		return nil, res.Error
	}
	return &schema, nil
}
