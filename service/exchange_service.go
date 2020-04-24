package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dexterorion/smartmei/errors"
	"github.com/dexterorion/smartmei/model"
)

// ExchangeService represents the exchange implementation service
type ExchangeService struct{}

// NewExchangeService creates new exchange service
func NewExchangeService() ExchangeService {
	return ExchangeService{}
}

// GetConversionRates returns a list of values converted from a currency to another ones
func (exs ExchangeService) GetConversionRates(ctx context.Context, from string, to []string) (map[string]float64, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	request, err := http.NewRequest("GET", "https://api.exchangeratesapi.io/latest?base=BRL&symbols=USD,EUR", nil)
	if err != nil {
		return nil, errors.ErrorBuildingRequest(err)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.ErrorDoingRequest(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.InvalidStatusCode(http.StatusOK, response.StatusCode)
	}
	byteBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrorProcessingResponseBody(err)
	}

	var exchangeData model.ExchangeResponse
	err = json.Unmarshal(byteBody, &exchangeData)
	if err != nil {
		return nil, errors.ErrorConvertingByteToMap(err)
	}

	return exchangeData.Rates, nil
}
