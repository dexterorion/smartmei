package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

	symbols := strings.Join(to, ",")
	url := fmt.Sprintf("https://api.exchangeratesapi.io/latest?base=%s&symbols=%s", from, symbols)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorBuildingRequest(err)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.ErrorDoingRequest(err)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

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
