package service

import (
	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/internal/repository"
)

type Urls interface {
	GetUrl(hash string) (string, error)
	GetUrlInfo(hash string) (models.UrlInfo, error)
	SetUrl(url models.SetUrl) (models.SetUrlResp, error)
	DeleteUrl(hash string) error
}

type Service struct {
	Urls
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Urls: NewUrlsService(repos.Urls),
	}
}
