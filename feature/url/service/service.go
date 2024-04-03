package service

import (
	"errors"
	"time"

	"github.com/RianIhsan/shorten_url/entities"
	"github.com/RianIhsan/shorten_url/feature/url"
	"github.com/RianIhsan/shorten_url/feature/url/dto"
	"github.com/RianIhsan/shorten_url/helper/cache"
	"github.com/RianIhsan/shorten_url/helper/random"
)

type urlService struct {
	readWrite url.URLRepositoryInterface
	cache     cache.RedisCache
}

func NewURLService(readWrite url.URLRepositoryInterface, cache cache.RedisCache) url.URLServiceInterface {
	return &urlService{
		readWrite: readWrite,
		cache: cache,
	}
}

func (s *urlService) CreateURL(req *dto.CreateURLRequest) (*entities.MstURL, error) {
	uniqShort, err := random.Generate()
	if err != nil {
		return nil, errors.New("failed generate unique code")
	}

	url := &entities.MstURL{
		OriginalURL: req.OriginalURL,
		ShorterURL:  uniqShort,
	}
	err = s.cache.SetRdsShortURL(uniqShort, 1*time.Hour, url.OriginalURL)
	if err != nil {
		return nil, err
	}

	data, err := s.readWrite.CreateURL(url)
	if err != nil {
		return nil, errors.New("failed save url data")
	}
	
	return data, nil
}

func (s *urlService) GetShortURL(url string) (*entities.MstURL, error) {
	data, err := s.readWrite.GetShortURL(url)
	if err != nil {
		return nil, errors.New("failed get short url")
	}
	return data, nil
}
