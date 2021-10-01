package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/models"
)

type Urls interface {
	GetUrl(hash string) string
	GetUrlInfo(hash string) models.Urls
	SetUrl(url models.SetUrl) models.SetUrlResp
	DeleteUrl(hash string)
}

type Repository struct {
	Urls
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Urls: NewUrlsMongo(client),
	}
}