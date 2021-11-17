package models

import (
	"time"
)

type Url struct {
	CreatedAt *time.Time   `bson:"created_at" json:"created-at,omitempty"`
	UpdatedAt *time.Time   `bson:"updated_at" json:"updated_at"`
	DeletedAt *time.Time   `bson:"deleted_at" json:"deleted_at"`
	Hash      string       `bson:"hash" json:"hash,omitempty"`
	Url       string       `bson:"url" json:"url,omitempty"`
	Count     int          `bson:"count" json:"count"`
	Access    bool         `bson:"access" json:"access"`
	Code      int          `bson:"code" json:"code"`
}

type SetUrl struct {
	Url string `json:"url,omitempty"`
}

type SetUrlResp struct {
	Long string     `json:"long"`
	Short string    `json:"short"`
}