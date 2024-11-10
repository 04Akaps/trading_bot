package job

import (
	"context"
	"github.com/04Akaps/trading_bot.git/common/http"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
	"log"
	"sync"
)

func (j *Job) currentPrice(c context.Context, cancel context.CancelFunc) {
	// TODO -> 가져올 금액 symbol 배열로 DB 조회
	symbols := map[string]bool{
		//"ETHBTC":  true,
		//"BNBBTC":  true,
		"POLUSDT": true,
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

				client := http.NewClient(t.APIHeaderKey, t.APIKey)

				err := client.GET(_currentPriceTimeTicker, emptyString, emptyString, &res)

				if err != nil {
					log.Println("Failed to get current price", "err", err)
					return
				}

				for _, o := range res {
					_, ok := symbols[o.Symbol]

					if ok {
						slackLoggerMap[k.ToString()][o.Symbol] = o.Price
					}
				}

			}

		}()

	}

	work.Wait()

	j.slackClient.CurrentPriceMessage(slackLoggerMap)

}
