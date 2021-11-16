package service

import (
	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/internal/repository"
)

type SrvUrls interface {
	GetUrl(hash string) (string, error)
	GetUrlInfo(hash string) (models.Url, error)
	SetUrl(url models.SetUrl) (models.SetUrlResp, error)
	DeleteUrl(hash string) error
}

type Service struct {
	SrvUrls
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		SrvUrls: NewUrlsService(repos.RepoUrls),
	}
}
