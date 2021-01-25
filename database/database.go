package database

import (
	"fmt"
	"net/url"

	_ "github.com/go-playground/validator/v10" // ...
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Config for db
type Config struct {
	NameDB   string
	Scheme   string
	User     string
	Password string
	Host     string
	Port     string
}

// Urls struct for db
type Urls struct {
	gorm.Model
	Hash   string `json:"hash"`
	URL    string `json:"url" validate:"required,url"`
	Count  int    `json:"count"`
	Access bool   `json:"access"`
	Code   int    `json:"code"`
}

// NewMSSQLDB conn with db
func NewMSSQLDB(cfg Config) (*gorm.DB, error) {
	query := url.Values{}
	query.Add("database", cfg.NameDB)
	u := &url.URL{
		Scheme:   cfg.Scheme,
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     fmt.Sprintf("%s%s", cfg.Host, cfg.Port),
		RawQuery: query.Encode(),
	}
	db, err := gorm.Open(sqlserver.Open(u.String()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&Urls{}); err != nil {
		return nil, err
	}

	return db, nil
}
