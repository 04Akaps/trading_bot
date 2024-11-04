package config

import (
	"github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
	_error "github.com/04Akaps/trading_bot.git/types/protocol/error"
)

type CryptoCurrency struct {
	APIHeaderKey string
	APIKey       string
	SecretKey    string
	Exchange     cryptoCurrency.Exchanger
}

func (c CryptoCurrency) IsSupportedExchange() {
	if _, ok := cryptoCurrency.SupportedExchanger[c.Exchange]; !ok {
		panic(_error.NOT_SUPPORTED_EXCNAHGER.E(string(c.Exchange)))
	}
}
