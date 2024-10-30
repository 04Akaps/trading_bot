package binance

import (
	"github.com/04Akaps/trading_bot.git/client/cryptoCurrency/impl"
	"github.com/04Akaps/trading_bot.git/config"
)

type Binance struct {
	cfg config.CryptoCurrency
}

func NewBinanceClient(cfg config.CryptoCurrency) impl.CurrencyClient {
	b := Binance{cfg: cfg}

	return b
}

func (c Binance) GetPrice(symbol string) string {
	//http.HttpClient.GET()
	return ""
}
