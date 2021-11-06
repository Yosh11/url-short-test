package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/models"
)

type Urls interface {
	GetUrl(hash string) (string, error)
	GetUrlInfo(hash string) (models.UrlInfo, error)
	SetUrl(url models.SetUrl) (models.SetUrlResp, error)
	DeleteUrl(hash string) error
}

type Repository struct {
	Urls
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Urls: NewUrlsMongo(client),
	}
}