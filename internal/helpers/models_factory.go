package helpers

import (
	"time"

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

// DeleteSafelyModel - Make model for delete method
func DeleteSafelyModel(time time.Time) bson.D {
	return bson.D{{"$set", bson.M{"deleted_at": time, "updated_at": time}}}
}

// IncrementCounterModel - Make model for redirect url counter
func IncrementCounterModel() bson.D {
	return bson.D{{"$inc", bson.M{"count": 1}}}
}

// UpdateTimeModel - Make model for redirect url counter
func UpdateTimeModel(time time.Time) bson.D {
	return bson.D{{"$set", bson.M{"updated_at": time}}}
}