package inspector

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/init/log"
	"github.com/Yosh11/url-short-test/internal/models"
)

type InsMongo struct {
	ctx  context.Context
	coll *mongo.Collection
}

var (
	wg = sync.WaitGroup{}
)

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
			wg.Add(1)
			go ins.check(res[count*i : count*(i+1)])
		}

		if len(res)%goCont != 0 {
			last := len(res) % goCont
			wg.Add(1)
			go ins.check(res[len(res)-last:])
		}

	} else {
		wg.Add(1)
		go ins.check(res)
	}

	wg.Wait()
	log.Info("Finish inspect all urls")
}

func (ins *InsMongo) check(urls []models.Url) {
	for _, v := range urls {
		var update bson.M
		res, err := http.Get(v.Url)
		if err != nil {
			update = bson.M{
				"deleted_at": time.Now().UTC(),
				"access":     false,
			}
		} else {
			update = bson.M{
				"updated_at": time.Now().UTC(),
				"code":       res.StatusCode,
				"access": func() bool {
					if res.StatusCode == 200 {
						return true
					}
					return false
				}(),
			}
		}

		mRes := ins.coll.FindOneAndUpdate(ins.ctx, bson.D{{"_id", v.Id}}, bson.D{{"$set", update}})
		log.CheckError(mRes.Err())
	}
	wg.Done()
}

func NewInsMongo(client *mongo.Client) *InsMongo {
	ctx := context.Background()
	coll := client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("COLLECTION"))
	return &InsMongo{ctx: ctx, coll: coll}
}
