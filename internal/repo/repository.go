package repo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Config struct for conn string
type Config struct {
	User string
	Pass string
	Host string
	Port string
}

// NewMongoDd init new repository
func NewMongoDd(ctx context.Context, c Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin&readPreference=primary&appname=MongoDB%%20Compass%%20Community&ssl=false",
			c.User, c.Pass, c.Host, c.Port),
	),
	)
	if err != nil {
		log.Fatal(err) // todo change to custom logger
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	} else {
		log.Println("DB connected")
	}
	return client, nil
}