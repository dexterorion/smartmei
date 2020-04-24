package dependency

import (
	"context"

	"github.com/dexterorion/smartmei/model"
	"github.com/dexterorion/smartmei/models"
)

// StartMeiService represents the smart mei service interface
type StartMeiService interface {
	CrawData(ctx context.Context, url string) (*model.DescriptionValueData, error)
}

// ExchangeService represents the exchange service interface
type ExchangeService interface {
	GetConversionRates(ctx context.Context, from string, to []string) (map[string]float64, error)
}

// Service interface defining this microservice's controller.
type Service interface {
	StartMeiService() StartMeiService
	ExchangeService() ExchangeService
	GetConvertedData(ctx context.Context, url, from string, to []string) (*models.CrawlerResponse, error)
}
