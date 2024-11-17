package job

import (
	"fmt"
	"github.com/04Akaps/trading_bot.git/common/http"
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
	"time"
)

func (j *Job) volumeTrace() {
	symbols := j.mongoDB.ScanTokenList()
	length := len(j.cfg.CryptoCurrency)

	slackLoggerMap := make(map[string]map[string]string, length)

	var work sync.WaitGroup
	work.Add(length)

	defer func() {
		j.volumeTraceInit.Store(false)
	}()

	for key, info := range j.cfg.CryptoCurrency {
		t := info
		k := key

		mapKey := k.ToString()

		if _, ok := slackLoggerMap[mapKey]; !ok {
			slackLoggerMap[mapKey] = make(map[string]string)
		}

		go func() {
			defer work.Done()

			switch k {
			case cryptoCurrency.Binance:

				client := http.NewClient(t.APIHeaderKey, t.APIKey)

				for s, _ := range symbols {

					var paramName []string
					var req []string

					if j.volumeTraceInit.Load() {
						paramName = []string{"symbol", "interval"}
						req = []string{s, "1d"}
					} else {
						paramName = []string{"symbol", "startTime", "interval"}
						now := time.Now().UTC()
						startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.UTC)
						startTime := fmt.Sprintf("%d", startOfDay.UnixMilli())
						req = []string{s, startTime, "1d"}
					}

					var res []interface{}

					err := client.GET(_traceVolume, paramName, req, &res)

					if err != nil {
						log.Println("Failed to get volume trace", "symbol", s, "err", err)
						return
					}

					models := make([]mongo.WriteModel, len(res))
					index := 0

					for _, datas := range res {
						data := datas.([]interface{})

						time := data[0]
						volume := data[5]

						models[index] = mongo.NewUpdateOneModel().
							SetFilter(bson.M{"time": time, "symbol": s}).
							SetUpdate(bson.M{"$set": bson.M{"time": time, "symbol": s, "volume": volume}}).
							SetUpsert(true)
						index++
					}

					if len(models) > 0 {
						j.mongoDB.UpdateBulk(models)
					}

					j.volumeUpdateChannel <- s
				}

			}

		}()

	}

	work.Wait()
}
