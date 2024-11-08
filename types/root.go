package types

type CurrentPriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type VolumeTicker struct {
	Symbol      string `json:"symbol"`
	PriceChange string `json:"priceChange"`
	HighPrice   string `json:"highPrice"`
	LowPrice    string `json:"lowPrice"`
	PairVolume  string `json:"volume"`
	QuoteVolume string `json:"quoteVolume"`
}

///futures/data/globalLongShortAccountRatio

//24h Volume(POL)
//159,613,810.00
//24h Volume(USDT)
//54,826,244.96

//{
//"symbol": "ETHBTC",
//"priceChange": "0.00195000",
//"priceChangePercent": "5.365",
//"weightedAvgPrice": "0.03769728",
//"prevClosePrice": "0.03635000",
//"lastPrice": "0.03830000",
//"lastQty": "12.93750000",
//"bidPrice": "0.03829000",
//"bidQty": "29.51970000",
//"askPrice": "0.03830000",
//"askQty": "10.33660000",
//"openPrice": "0.03635000",
//"highPrice": "0.03875000",
//"lowPrice": "0.03635000",
//"volume": "85947.91270000",
//"quoteVolume": "3240.00237955",
//"openTime": 1730941432501,
//"closeTime": 1731027832501,
//"firstId": 471985580,
//"lastId": 472279634,
//"count": 294055
//},
