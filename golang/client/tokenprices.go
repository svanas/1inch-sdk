package client

import (
	"context"
	"fmt"
	"net/http"

	"dev-portal-sdk-go/client/tokenprices"
)

type PricesResponse map[string]string

func (c Client) GetTokenPrices(params tokenprices.PricesParameters) (string, *http.Response, error) {
	// TODO accept context
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/price/v1.1/1", c.BaseURL), nil)

	err = validateParameters(params)
	if err != nil {
		return "", nil, err
	}

	switch params.Currency {
	case "":
	case "WEI":
	default:
		req.URL.RawQuery += fmt.Sprintf("currency=%s", params.Currency)
	}
	if err != nil {
		return "", nil, err
	}

	var exStr PricesResponse
	res, err := c.Do(context.Background(), req, &exStr)
	if err != nil {
		return "", nil, err
	}

	return fmt.Sprintf("Pirces: %v", exStr), res, nil
}

func validateParameters(params tokenprices.PricesParameters) error {
	if !contains(tokenprices.CurrencyTypeValues, params.Currency) {
		return fmt.Errorf("currency value %s is not valid", params.Currency)
	}
	return nil
}

// Make a helpers class
func contains(slice []tokenprices.CurrencyType, item tokenprices.CurrencyType) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
