package service

import (
	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/internal/repository"
)

type UrlsService struct {
	repo repository.RepoUrls
}

func NewUrlsService(repo repository.RepoUrls) *UrlsService {
	return &UrlsService{repo: repo}
}

func (u *UrlsService) GetUrl(hash string) (string, error) {
	return u.repo.GetUrl(hash)
}

func (u *UrlsService) GetUrlInfo(hash string) (models.Url, error) {
	return u.repo.GetUrlInfo(hash)
}

func (u *UrlsService) SetUrl(url models.SetUrl) (models.SetUrlResp, error) {
	return u.repo.SetUrl(url)
}

func (u *UrlsService) DeleteUrl(hash string) error {
	return u.repo.DeleteUrl(hash)
}
