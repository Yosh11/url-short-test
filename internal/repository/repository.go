package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/models"
)

type RepoUrls interface {
	Create(url models.SetUrl) (models.SetUrlResp, error)
	Get(hash string) (models.Url, error)
	Update(id primitive.ObjectID, newData bson.D) (models.Url, error)
	Delete(hash string) error
}

type Repository struct {
	RepoUrls
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		RepoUrls: NewUrlsMongo(client),
	}
}