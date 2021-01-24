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
	Port     uint32
}

// Urls struct for db
type Urls struct {
	gorm.Model
	ID    string `json:"id"`
	URL   string `json:"url" validate:"required,url"`
	Count int    `json:"count"`
}

// NewMSSQLDB conn with db
func NewMSSQLDB(cfg Config) (*gorm.DB, error) {
	query := url.Values{}
	query.Add("database", "url-short")
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword("sa", "p2ompiTJ"),
		Host:     fmt.Sprintf("%s:%d", "localhost", 1433),
		RawQuery: query.Encode(),
	}
	db, err := gorm.Open(sqlserver.Open(u.String()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
