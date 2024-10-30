package currencyapi_test

import (
	"fmt"
	"github.com/Thanasak1412/http-request/currencyapi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRate(t *testing.T) {
	t.Run("Empty currency code", func(t *testing.T) {
		_, err := currencyapi.GetRate("")

		if err != nil {
			t.Error("expected an error for empty currency code, got gone")
		}
	})

	t.Run("Invalid currency code", func(t *testing.T) {
		data, err := currencyapi.GetRate("INVALID")

		if err != nil || data == nil {
			t.Error("Invalid current code should return error")
		}
	})

	t.Run("Valid currency code", func(t *testing.T) {
		//	Mock HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"last": "1234.56"}`)
		}),
		)

		defer server.Close()

		// Replace the apiUrl with mock server URL
		originalUrl := currencyapi.ApiUrl
		currencyapi.ApiUrl = server.URL + "/api/ticker/%s/USD"

		defer func() {
			currencyapi.ApiUrl = originalUrl
		}()

		rate, err := currencyapi.GetRate("BTC")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		expectedPrice := 1234.56
		if rate.Price != expectedPrice {
			t.Errorf("expected price %v, got %v", expectedPrice, rate.Price)
		}

	})
}
