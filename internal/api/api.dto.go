package api

type HealthResponse struct {
	Message string `json:"message"`
}

type CrawlResponse struct {
	Data string `json:"data"`
}

type Currency struct {
	Usd float32 `json:"usd"`
}

/*
	{
		"dai": {
			"usd": 0.999821
			},
		"paypal-usd": {
			"usd": 0.999818
		},
		"tether": {
			"usd": 0.999432
		},
		"usd-coin": {
			"usd": 0.999701
		},
		"usds": {
			"usd": 0.999776
		}
	}
*/
type CoinGeckoApiResponse struct {
	Dai       Currency `json:"dai"`
	PaypalUsd Currency `json:"paypal-usd"`
	Tether    Currency `json:"tether"`
	UsdCoin   Currency `json:"usd-coin"`
	Usds      Currency `json:"usds"`
}

type SseStatsResponse struct {
	Data CoinGeckoApiResponse `json:"data"`
}
