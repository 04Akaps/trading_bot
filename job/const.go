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
	// 10월 10일 이후의 데이터만 수집
	_leastStartTime = "1728518400000" // 10월 10일 오전 9시
	_traceVolume    = "https://api.binance.com/api/v3/klines"
)
