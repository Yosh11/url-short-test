package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/models"
)

type RepoUrls interface {
	GetUrl(hash string) (string, error)
	GetUrlInfo(hash string) (models.Url, error)
	SetUrl(url models.SetUrl) (models.SetUrlResp, error)
	DeleteUrl(hash string) error
}

type Repository struct {
	RepoUrls
}

func NewRepository(ctx context.Context, client *mongo.Client) *Repository {
	return &Repository{
		RepoUrls: NewUrlsMongo(ctx, client),
	}
}