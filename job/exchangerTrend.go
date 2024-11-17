package job

import (
	"fmt"
	"github.com/04Akaps/trading_bot.git/common/http"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
	"log"
	"sort"
	"strconv"
	"sync"
	"time"
)

func (j *Job) exchangerTrend() {
	length := len(j.cfg.CryptoCurrency)

	var work sync.WaitGroup
	work.Add(length)

	j.slackClient.Top5VolumeDiffStarter()

	for key, info := range j.cfg.CryptoCurrency {
		t := info
		k := key

		switch k {
		case cryptoCurrency.Binance:
			client := http.NewClient(t.APIHeaderKey, t.APIKey)
			var result []types.Top5VolumeDiff

			for s, _ := range j.symbols {
				kst, _ := time.LoadLocation("Asia/Seoul")
				now := time.Now().UTC()
				yesterday := now.AddDate(0, 0, -1)
				startOfDayKST := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 9, 0, 0, 0, kst)
				yesterDay := fmt.Sprintf("%d", startOfDayKST.UnixMilli())

				req := []string{s, yesterDay, "1d"}
				paramName := []string{"symbol", "startTime", "interval"}
				var res []interface{}

				err := client.GET(_traceVolume, paramName, req, &res)

				if err != nil {
					log.Println("Failed to get volume trace", "symbol", s, "err", err)
					return
				}

				var beforeVolume float64
				var currentVolume float64

				if len(res) == 2 {
					// 최소 어제치 거래까지 있는 애들
					for i, datas := range res {
						data := datas.([]interface{})
						volume := data[5].(string)
						fVolume, _ := strconv.ParseFloat(volume, 64)

						if i == 0 {
							// 어제 거래
							beforeVolume = fVolume
						} else {
							// 오늘 거래
							currentVolume = fVolume
						}

					}
				}

				percentChange := 0.0
				if beforeVolume > 0 {
					percentChange = ((currentVolume - beforeVolume) / beforeVolume) * 100
				}

				// 결과 저장
				result = append(result, types.Top5VolumeDiff{
					Symbol:        s,
					BeforeVolume:  beforeVolume,
					CurrentVolume: currentVolume,
					Diff:          percentChange,
				})

			}

			sort.Slice(result, func(i, j int) bool {
				return result[i].Diff > result[j].Diff
			})

			// 상위 5개 데이터 출력
			top5 := result
			if len(result) > 5 {
				top5 = result[:5]
			}

			j.slackClient.Top5VolumeDiffTrend(top5)
		}

	}

}
