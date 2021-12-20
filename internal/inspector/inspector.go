package inspector

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/models"
)

type InsUrls interface {
	StartInspect()
	check(urls []models.Url)
}

type Inspector struct {
	InsUrls
}

func NewInspector(client *mongo.Client) *Inspector {
	return &Inspector{
		InsUrls: NewInsMongo(client),
	}
}
