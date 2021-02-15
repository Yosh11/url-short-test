package inspector

import (
    "log"
	"net/http"
	"sync"
	"time"

	"github.com/Yosh11/url-short-test/database"
	"github.com/Yosh11/url-short-test/pkg/handlers"
)

var (
	wg sync.WaitGroup
	db = handlers.InitDB()
)

// Check access to the source
func Check() {
	var count, i int64
	var table database.Urls
	var urls []string

	db.Model(&table).Distinct("hash").Count(&count)
	db.Raw("SELECT url FROM urls").Scan(&urls)

	log.Println("[Inspector] Status сheck")

	for i = 0; i < count; i++ {
		wg.Add(1)
		go runCheck(urls[i], &wg)
	}
	wg.Wait()
}

func runCheck(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	var pattern database.Urls
	db.Where("url = ?", url).Find(&pattern)
	time.Sleep(1 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%s -> %s", url, err.Error())
		pattern.Access = false
		pattern.Code = 409
		db.Save(&pattern)
	} else {
		log.Printf("%s -> %s", url, resp.Status)
		if resp.StatusCode == 404 {
			pattern.Code = resp.StatusCode
			pattern.Access = false
			db.Save(&pattern)
		} else {
			if resp.StatusCode != pattern.Code {
				pattern.Code = resp.StatusCode
				db.Save(&pattern)
			}
		}
	}
}
