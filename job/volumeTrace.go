package job

import (
	"context"
	"fmt"
	"github.com/04Akaps/trading_bot.git/common/http"
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
	"log"
	"sync"
	"time"
)

func (j *Job) volumeTrace(c context.Context, cancel context.CancelFunc) {
	// 5 번쨰가 volume
	symbols := j.mongoDB.ScanTokenList()
	length := len(j.cfg.CryptoCurrency)

	slackLoggerMap := make(map[string]map[string]string, length)

	var work sync.WaitGroup
	work.Add(length)

	defer j.volumeTraceInit.Store(false)

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
				var startTime string

				if j.volumeTraceInit.Load() {
					startTime = _leastStartTime
				} else {
					now := time.Now().UTC() // 현재 UTC 시간
					startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.UTC)
					startTime = fmt.Sprintf("%d", startOfDay.UnixMilli()) // 밀리초 단위로 변환
				}

				client := http.NewClient(t.APIHeaderKey, t.APIKey)

				for s, _ := range symbols {

					var res []interface{}

					err := client.GET(
						_traceVolume,
						[]string{"symbol", "startTime", "interval"},
						[]string{s, startTime, "1d"},
						&res,
					)

					if err != nil {
						log.Println("Failed to get volume trace", "symbol", s, "err", err)
						return
					}

					fmt.Println(res)

					//[
					//	1731196800000,
					//	"0.38550000",
					//	"0.41090000",
					//	"0.38110000",
					//	"0.39790000",
					//	"60232608.20000000",
					//	1731283199999,
					//	"23813569.16730000",
					//	78243,
					//	"30583249.40000000",
					//	"12097183.77477000",
					//	"0"
					//]
				}

			}

		}()

	}

	work.Wait()
	//
	//j.slackClient.CurrentPriceMessage(slackLoggerMap)

}
