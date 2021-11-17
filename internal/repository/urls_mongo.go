package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/helpers"
	"github.com/Yosh11/url-short-test/internal/models"
	"github.com/Yosh11/url-short-test/tools/genhash"
)

type UrlsMongo struct {
	ctx context.Context
	coll *mongo.Collection
}


func NewUrlsMongo(client *mongo.Client) *UrlsMongo {
	ctx := context.Background()
	coll := client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("COLLECTION"))
	return &UrlsMongo{ctx: ctx, coll: coll}
}

func (u *UrlsMongo) Get(hash string) (models.Url, error) {
	var ret models.Url

	result := u.coll.FindOne(u.ctx, bson.D{{"hash", hash}}); if result.Err() != nil {
		return models.Url{}, result.Err()
	}

	err := result.Decode(&ret); if err != nil {
		return models.Url{}, err
	}

	return ret, nil
}

func (u *UrlsMongo) Create(url models.SetUrl) (models.SetUrlResp, error) {
	timeNow := time.Now().UTC()
	hash := genhash.Generate()

	newUrl := models.Url{
		CreatedAt: &timeNow,
		UpdatedAt: nil,
		DeletedAt: nil,
		Hash:      hash,
		Url:       url.Url,
		Count:     0,
		Access:    true,
		Code:      0,
	}

	newUrlDB := helpers.CreateUrlFactory(newUrl)

	_, err := u.coll.InsertOne(u.ctx, newUrlDB); if err != nil {
		return models.SetUrlResp{}, err
	}

	ret := models.SetUrlResp{
		Long:  url.Url,
		Short: fmt.Sprintf("http://%s:%s/urls/%s",os.Getenv("HOST_API"), os.Getenv("PORT_API"), hash),
	}

	return ret, nil
}

func (u *UrlsMongo) Update(id primitive.ObjectID, newData bson.D) (models.Url, error) {
	_, err := u.coll.UpdateByID(u.ctx, id, newData); if err != nil {
		return models.Url{}, err
	}
	return models.Url{}, nil
}

func (u *UrlsMongo) Delete(hash string) error {
	_, err := u.coll.DeleteOne(u.ctx, bson.D{{"hash", hash}}); if err != nil {
		return err
	}
	return nil
}


