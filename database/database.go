package database

import (
    "fmt"

    _ "github.com/go-playground/validator/v10" // ...
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Config for db
type Config struct {
	NameDB   string
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

// NewDB conn with db
func NewDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
	    cfg.Host, cfg.User, cfg.Password, cfg.NameDB, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&Urls{}); err != nil {
		return nil, err
	}

	return db, nil
}
