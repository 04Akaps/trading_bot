package types

type CurrentPriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type VolumeTicker struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	OpenPrice          string `json:"openPrice"`
}

type TradingDayTicker struct {
	Symbol      string `json:"symbol"`
	QuoteVolume string `json:"quoteVolume"`
	Volume      string `json:"volume"`
}

type VolumeTrend struct {
	VolumeTicker
	TradingDayTicker
}
