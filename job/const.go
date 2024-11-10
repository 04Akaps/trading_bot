package job

var (
	emptyString = []string{}
)

const (
	// TODO changer 마다 url 변경
	_currentPriceTimeTicker = "https://api1.binance.com/api/v3/ticker/price"
)

const (
	// TODO changer 마다 url 변경
	_currentVolumeTimeTicker = "https://api1.binance.com/api/v3/ticker/24hr"
	_tradingDay              = "https://api1.binance.com/api/v3/ticker/tradingDay"
)

const (
	_traceVolume = "https://api.binance.com/api/v3/klines"
)
