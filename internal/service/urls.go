package service

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Yosh11/url-short-test/internal/helpers"
	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/internal/repository"
)

type UrlsService struct {
	repo repository.RepoUrls
}

func NewUrlsService(repo repository.RepoUrls) *UrlsService {
	return &UrlsService{repo: repo}
}

func (u *UrlsService) GetUrl(hash string) (models.Url, error) {
	obj, err := u.repo.Get(hash)
	if err != nil {
		return models.Url{}, err
	}

	if obj.DeletedAt != nil {
		return models.Url{}, errors.New("URL has been removed")
	}

	if !obj.Access {
		return models.Url{}, errors.New("the URL is currently unavailable")
	}

	_, err = u.repo.Update(obj.Id, helpers.IncrementCounterModel())
	if err != nil {
		return models.Url{}, err
	}

	_, err = u.repo.Update(obj.Id, helpers.UpdateTimeModel(time.Now().UTC()))
	if err != nil {
		return models.Url{}, err
	}

	return obj, nil
}

func (u *UrlsService) GetUrlInfo(hash string) (models.Url, error) {
	return u.repo.Get(hash)
}

func (u *UrlsService) SetUrl(url models.SetUrl) (models.SetUrlResp, error) {
	return u.repo.Create(url)
}

func (u *UrlsService) DeleteUrl(id primitive.ObjectID) error {
	return u.repo.Delete(id)
}
