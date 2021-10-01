package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/models"
)

type UrlsMongo struct {
	coll *mongo.Collection
}

func NewUrlsMongo(client *mongo.Client) *UrlsMongo {
	coll := client.Database(database).Collection(collectionUrls)
	return &UrlsMongo{coll: coll}
}

func (u *UrlsMongo) GetUrl(hash string) string {
	panic("implement me")
}

func (u *UrlsMongo) GetUrlInfo(hash string) models.Urls {
	panic("implement me")
}

func (u *UrlsMongo) SetUrl(url models.SetUrl) models.SetUrlResp {
	panic("implement me")
}

func (u *UrlsMongo) DeleteUrl(hash string) {
	panic("implement me")
}


