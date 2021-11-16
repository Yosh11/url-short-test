package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/tools/genhash"
)

type UrlsMongo struct {
	ctx context.Context
	coll *mongo.Collection
}

func NewUrlsMongo(ctx context.Context, client *mongo.Client) *UrlsMongo {
	coll := client.Database(database).Collection(collectionUrls)
	return &UrlsMongo{ctx: ctx, coll: coll}
}

func (u *UrlsMongo) GetUrl(hash string) (string, error) {
	return "", nil
}

func (u *UrlsMongo) GetUrlInfo(hash string) (models.Url, error) {
	panic("implement me")
}

func (u *UrlsMongo) SetUrl(url models.SetUrl) (models.SetUrlResp, error) {
	timeNow := time.Now().UTC()
	hash := genhash.Generate()

	newUrl := models.Url{
		ID:        primitive.ObjectID{},
		CreatedAt: &timeNow,
		UpdatedAt: nil,
		DeletedAt: nil,
		Hash:      hash,
		Url:       url.Url,
		Count:     0,
		Access:    true,
		Code:      0,
	}

	_, err := u.coll.InsertOne(u.ctx, newUrl); if err != nil {
		return models.SetUrlResp{}, err
	}

	ret := models.SetUrlResp{
		Long:  url.Url,
		Short: fmt.Sprintf("%s:%s/%s",os.Getenv("HOST_API"), os.Getenv("PORT_API"), hash),
	}

	return ret, nil
}

func (u *UrlsMongo) DeleteUrl(hash string) error {
	panic("implement me")
}


