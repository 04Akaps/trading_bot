package job

import (
	"github.com/04Akaps/trading_bot.git/common/http"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
)

func (j *Job) volumeTraceDiffChecker() {
	for {
		select {
		case symbol := <-j.volumeUpdateChannel:
			go func() {
				total, currentVolume, diff := j.mongoDB.GetVolumeInfo(symbol)

				j.slackClient.VolumeTracker(symbol, total, currentVolume, diff)
			}()
		}
	}
}

func (j *Job) binanceAllSymbols() map[string]bool {
	t := j.cfg.CryptoCurrency[cryptoCurrency.Binance]

	client := http.NewClient(t.APIHeaderKey, t.APIKey)

	var mapperRes []*types.SymbolMapper

	err := client.GET(_currentVolumeTimeTicker, emptyString, emptyString, &mapperRes)

	if err != nil {
		panic("Failed to get binance All Symbols")
	}

	symbols := make(map[string]bool, len(mapperRes))

	for _, s := range mapperRes {
		symbols[s.Symbol] = true
	}

	return symbols

}
