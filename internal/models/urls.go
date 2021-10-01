package models

import "time"

type Urls struct {
	CreatedAt *time.Time   `json:"created-at"`
	UpdatedAt *time.Time   `json:"updated_at"`
	DeletedAt *time.Time   `json:"deleted_at"`
	Hash      string       `json:"hash"`
	Url       string       `json:"url"`
	Count     int          `json:"count"`
	Access    bool         `json:"access"`
	Code      int          `json:"code"`
}

type SetUrl struct {
	Url string `json:"url"`
}

type SetUrlResp struct {
	Long string     `json:"long"`
	Short string    `json:"short"`
}