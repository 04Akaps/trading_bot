package cryptoCurrency

type Exchanger string

const (
	Binance = Exchanger("binance")
)

var SupportedExchanger = map[Exchanger]bool{
	Binance: true,
}

func (e Exchanger) ToString() string {
	return string(e)
}
