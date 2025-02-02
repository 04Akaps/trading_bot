package job

import (
	"fmt"
	"github.com/04Akaps/trading_bot.git/common/http"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
	"log"
	"strings"
	"sync"
)

func (j *Job) volumeTrend() {
	symbols := j.scanSymbols

	length := len(j.cfg.CryptoCurrency)

	slackLoggerMap := make(map[string]map[string]types.VolumeTrend, length)

	var work sync.WaitGroup
	work.Add(length)

	for key, info := range j.cfg.CryptoCurrency {

		t := info
		k := key

		mapKey := k.ToString()

		if _, ok := slackLoggerMap[mapKey]; !ok {
			slackLoggerMap[mapKey] = make(map[string]types.VolumeTrend)
		}

		go func() {
			defer work.Done()

			switch k {
			case cryptoCurrency.Binance:

				client := http.NewClient(t.APIHeaderKey, t.APIKey)

				var volumeRes []*types.VolumeTicker

				err := client.GET(_currentVolumeTimeTicker, emptyString, emptyString, &volumeRes)

				if err != nil {
					log.Println("Failed to get current volume", "err", err)
					return
				}

				for _, o := range volumeRes {
					_, ok := symbols[o.Symbol]

					if ok {
						slackLoggerMap[mapKey][o.Symbol] = types.VolumeTrend{
							VolumeTicker: types.VolumeTicker{
								PriceChange:        o.PriceChange,
								PriceChangePercent: o.PriceChangePercent,
								HighPrice:          o.HighPrice,
								LowPrice:           o.LowPrice,
								OpenPrice:          o.OpenPrice,
							},
						}
					}
				}

				var tradingRes []*types.TradingDayTicker

				var quotedSymbols []string

				for symbol, _ := range symbols {
					quotedSymbols = append(quotedSymbols, fmt.Sprintf("\"%s\"", symbol))
				}

				req := fmt.Sprintf("[%s]", strings.Join(quotedSymbols, ","))

				err = client.GET(_tradingDay, []string{"symbols"}, []string{req}, &tradingRes)

				if err != nil {
					log.Println("Failed to get trading day", "err", err)
					return
				}

				for _, o := range tradingRes {
					ticker := slackLoggerMap[mapKey][o.Symbol]
					ticker.TradingDayTicker = types.TradingDayTicker{
						QuoteVolume: o.QuoteVolume,
						Volume:      o.Volume,
					}

					slackLoggerMap[mapKey][o.Symbol] = ticker
				}

			}

		}()

	}

	work.Wait()

	j.slackClient.VolumeMessage(slackLoggerMap)
}
