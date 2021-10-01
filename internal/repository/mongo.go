package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Yosh11/url-short-test/init/err"
)

const (
	database = "url_short_test"
	collectionUrls = "urls"
)

// Config struct for conn string
type Config struct {
	User string
	Pass string
	Host string
	Port string
}

func InitMongo() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c := Config{
		User: os.Getenv("USER"),
		Pass: os.Getenv("PASS"),
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT_DB"),
	}
	client, e := mongo.NewClient(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin&readPreference=primary&appname=MongoDB%%20Compass%%20Community&ssl=false",
			c.User, c.Pass, c.Host, c.Port),
		),
	)
	err.CheckFatal(e, "Problem with NewClient for Mongo")
	e = client.Connect(ctx)
	err.CheckFatal(e, "Problem with connect to Mongo")

	if e = client.Ping(ctx, readpref.Primary()); e != nil {
		err.Fata(e, "Fail DB conn")
	} else {
		log.Println("DB connected")
	}
	return client
}
