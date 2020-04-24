package service

import (
	"context"
	"time"

	"github.com/dexterorion/smartmei/model"

	"github.com/dexterorion/smartmei/models"

	"github.com/dexterorion/smartmei/errors"

	"github.com/dexterorion/smartmei/dependency"
)

// Service represents the service structure that holds other controllers.
type Service struct {
	smartmeiService dependency.StartMeiService
	exchangeService dependency.ExchangeService
}

// New creates new service struct.
func New() Service {
	return Service{
		smartmeiService: NewStartMeiService(),
		exchangeService: NewExchangeService(),
	}
}

// StartMeiService returns reference to smartmei service
func (s Service) StartMeiService() dependency.StartMeiService {
	return s.smartmeiService
}

// ExchangeService returns reference to exchange service
func (s Service) ExchangeService() dependency.ExchangeService {
	return s.exchangeService
}

// GetConvertedData craw data and then return the conversion and more infos
func (s Service) GetConvertedData(ctx context.Context, url, from string, to []string) (*models.CrawlerResponse, error) {
	timeRequest := time.Now().Format(time.RFC3339)

	crawledData, err := s.StartMeiService().CrawData(ctx, url)
	if err != nil {
		return nil, errors.ErrorCrawlingData(err)
	}
	rates, err := s.ExchangeService().GetConversionRates(ctx, from, to)
	if err != nil {
		return nil, errors.ErrorGettingExchange(err)
	}

	response := &models.CrawlerResponse{
		Date:        timeRequest,
		Description: crawledData.Description,
		Rates: models.CrawlerResponseRates{
			Brl: crawledData.Value,
			Usd: crawledData.Value * rates[model.CurrencyUSD],
			Eur: crawledData.Value * rates[model.CurrencyEUR],
		},
	}

	return response, nil
}
