package types

type CurrentPriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type SymbolMapper struct {
	Symbol string `json:"symbol"`
}

type Top5VolumeDiff struct {
	CurrentVolume float64
	BeforeVolume  float64
	Diff          float64
	Symbol        string
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

type VolumeDocument struct {
	Time   int64  `json:"time" bson:"time"`
	Symbol string `json:"symbol" bson:"symbol"`
	Volume string `json:"volume" bson:"volume"`
}
