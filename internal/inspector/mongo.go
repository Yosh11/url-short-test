package inspector

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/init/log"
	"github.com/Yosh11/url-short-test/internal/models"
)

type InsMongo struct {
	ctx  context.Context
	coll *mongo.Collection
}

func (ins *InsMongo) StartInspect() {
	log.Info("Start inspect all urls")

	var res []models.Url
	goCont, err := strconv.Atoi(os.Getenv("DEFAULT_COUNT_GOROUTINES"))
	log.CheckFatal(err)

	cursor, err := ins.coll.Find(ins.ctx, bson.M{})
	log.CheckError(err)

	err = cursor.All(ins.ctx, &res)
	log.CheckError(err)

	if len(res) >= 10 {
		count := len(res) / goCont
		for i := 0; i < goCont; i++ {
			go ins.check(res[count*i : count*(i+1)])
		}

		if len(res)%goCont != 0 {
			last := len(res) % goCont
			go ins.check(res[len(res)-last:])
		}

	} else {
		go ins.check(res)
	}
}

func (ins *InsMongo) check(urls []models.Url) {
	for _, v := range urls {
		_, err := http.Get(v.Url)
		log.CheckError(err)
	}
}

func NewInsMongo(client *mongo.Client) *InsMongo {
	ctx := context.Background()
	coll := client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("COLLECTION"))
	return &InsMongo{ctx: ctx, coll: coll}
}
