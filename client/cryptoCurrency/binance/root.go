package binance

import (
	"github.com/04Akaps/trading_bot.git/client/cryptoCurrency/impl"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/adshao/go-binance/v2"
)

type Binance struct {
	cfg    config.CryptoCurrency
	client *binance.Client
}

func NewBinanceClient(cfg config.CryptoCurrency) impl.CurrencyClient {
	b := Binance{cfg: cfg}

	b.client = binance.NewClient(cfg.APIKey, cfg.SecretKey)

	//b.client.
	return b
}
