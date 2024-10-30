package impl

type CurrencyClient interface {
	GetPrice(symbol string) string
}
