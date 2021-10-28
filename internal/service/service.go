package service

import (
	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/internal/repository"
)

type Urls interface {
	GetUrl(hash string) string
	GetUrlInfo(hash string) models.Urls
	SetUrl(url models.SetUrl) models.SetUrlResp
	DeleteUrl(hash string)
}

type Service struct {
	Urls
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Urls: NewUrlsService(repos.Urls),
	}
}
