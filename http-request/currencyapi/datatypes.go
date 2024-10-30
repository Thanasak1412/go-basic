package currencyapi

type Rate struct {
	CurrencyCode string
	Price        float64 `json:"last"`
}
