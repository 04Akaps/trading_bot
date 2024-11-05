package job

import (
	"context"
	"fmt"
	"github.com/04Akaps/trading_bot.git/common/http"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
	"log"
	"sync"
)

const (
	// TODO changer 마다 url 변경
	_currentPriceTimeTicker = "https://api1.binance.com/api/v3/ticker/price"
)

func (j *Job) CurrentPrice(c context.Context, cancel context.CancelFunc) {
	// TODO -> 가져올 금액 symbol 배열로 DB 조회
	symbols := map[string]bool{
		"ETHBTC": true,
		"BNBBTC": true,
	}

	length := len(j.cfg.CryptoCurrency)

	slackLoggerMap := make(map[string]map[string]string, length)

	var work sync.WaitGroup

	work.Add(length)

	for key, info := range j.cfg.CryptoCurrency {

		t := info
		k := key

		if _, ok := slackLoggerMap[k.ToString()]; !ok {
			slackLoggerMap[k.ToString()] = make(map[string]string)
		}

		go func() {
			defer work.Done()

			switch k {
			case cryptoCurrency.Binance:
				var res []*types.CurrentPriceTicker

				err := http.HttpClient.GetCurrentPriceTicker(_currentPriceTimeTicker, t.APIHeaderKey, t.APIKey, &res)

				if err != nil {
					log.Println("Failed to get current price", "err", err)
					return
				}

				for _, o := range res {
					fmt.Println(o.Symbol)
					_, ok := symbols[o.Symbol]

					if ok {
						slackLoggerMap[cryptoCurrency.Binance.ToString()][o.Symbol] = o.Price
					}
				}

			}

		}()

	}

	work.Wait()

	for k, v := range slackLoggerMap {
		fmt.Println(k, v)
	}

}
