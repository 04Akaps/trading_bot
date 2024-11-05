package cryptoCurrency

import (
	"github.com/04Akaps/trading_bot.git/client/cryptoCurrency/binance"
	"github.com/04Akaps/trading_bot.git/client/cryptoCurrency/impl"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
)

var constructor = map[cryptoCurrency.Exchanger]func(cfg config.CryptoCurrency) impl.CurrencyClient{
	cryptoCurrency.Binance: binance.NewBinanceClient,
}

type CryptoCurrency struct {
	exchanger map[cryptoCurrency.Exchanger]impl.CurrencyClient
}

func NewCryptoCurrency(cfg map[cryptoCurrency.Exchanger]config.CryptoCurrency) CryptoCurrency {
	c := CryptoCurrency{
		exchanger: make(map[cryptoCurrency.Exchanger]impl.CurrencyClient, len(cfg)),
	}

	for k, v := range cfg {
		newClient := constructor[k]   // 이 부분이 함수로 할당된 것을 확인
		c.exchanger[k] = newClient(v) // 여기서 v를 인자로 전달
	}

	return c
}

func (c CryptoCurrency) GetTokenPrice(cex cryptoCurrency.Exchanger, symbol string) string {
	return c.exchanger[cex].GetPrice(symbol)
}
