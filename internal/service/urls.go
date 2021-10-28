package service

import (
	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/internal/repository"
)

type UrlsService struct {
	repo repository.Urls
}

func (u UrlsService) GetUrl(hash string) string {
	panic("implement me")
}

func (u UrlsService) GetUrlInfo(hash string) models.Urls {
	panic("implement me")
}

func (u UrlsService) SetUrl(url models.SetUrl) models.SetUrlResp {
	panic("implement me")
}

func (u UrlsService) DeleteUrl(hash string) {
	panic("implement me")
}

func NewUrlsService(repo repository.Urls) *UrlsService {
	return &UrlsService{repo: repo}
}
