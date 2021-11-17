package helpers

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Yosh11/url-short-test/internal/models"
)

// CreateUrlFactory - Make db format
func CreateUrlFactory(url models.Url) bson.M {
	return bson.M{
		"created_at": url.CreatedAt,
		"updated_at": url.UpdatedAt,
		"deleted_at": url.DeletedAt,
		"hash": url.Hash,
		"url": url.Url,
		"count": url.Count,
		"access": url.Access,
		"code": url.Code,
	}
}
