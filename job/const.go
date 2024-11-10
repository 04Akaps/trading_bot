package job

var (
	emptyString = []string{}
)

// TODO changer 마다 url 변경 DB로 관리

const (
	_currentPriceTimeTicker = "https://api1.binance.com/api/v3/ticker/price"
)

const (
	_currentVolumeTimeTicker = "https://api1.binance.com/api/v3/ticker/24hr"
	_tradingDay              = "https://api1.binance.com/api/v3/ticker/tradingDay"
)

const (
	_traceVolume = "https://api.binance.com/api/v3/klines"
)
