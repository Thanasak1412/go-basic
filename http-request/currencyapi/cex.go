package currencyapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var ApiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currencyCode string) (*Rate, error) {
	if currencyCode == "" {
		return nil, errors.New("no currency code")
	}

	resp, err := http.Get(fmt.Sprintf(ApiUrl, currencyCode))
	rate := Rate{}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		var result Response
		err = json.Unmarshal(bodyBytes, &result)

		if err != nil {
			return nil, err
		}

		rate.CurrencyCode = currencyCode

		price, _ := strconv.ParseFloat(result.Last, 64)

		rate.Price = price
	} else {
		return nil, errors.New(resp.Status)
	}

	return &rate, nil
}
