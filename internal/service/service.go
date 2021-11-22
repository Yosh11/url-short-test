package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/internal/repository"
)

type SrvUrls interface {
	GetUrl(hash string) (models.Url, error)
	GetUrlInfo(hash string) (models.Url, error)
	SetUrl(url models.SetUrl) (models.SetUrlResp, error)
	DeleteUrl(hash primitive.ObjectID) error
}

type Service struct {
	SrvUrls
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		SrvUrls: NewUrlsService(repos.RepoUrls),
	}
}
