package service

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dexterorion/smartmei/errors"

	"github.com/dexterorion/smartmei/model"
)

// SmartMeiService represents the smart mei implementation service
type SmartMeiService struct{}

// NewStartMeiService creates new smartmei service
func NewStartMeiService() SmartMeiService {
	return SmartMeiService{}
}

// CrawData gets required data from url website
func (sms SmartMeiService) CrawData(ctx context.Context, url string) (*model.DescriptionValueData, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
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

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrorProcessingResponseBody(err)
	}

	sts := strings.Split(string(bodyBytes), "<div class=\"text-center col-sm-4 col-xs-6 tarifas-2-3-2\">")

	if len(sts) != 2 {
		return nil, errors.InvalidArraySize(2, len(sts))
	}

	foundValue := strings.Split(sts[1], "</div>")
	description := strings.TrimSpace(foundValue[0])
	value := strings.Split(description, " ")

	if len(sts) != 2 {
		return nil, errors.InvalidArraySize(2, len(value))
	}

	val, err := strconv.ParseFloat(strings.Replace(value[1], ",", ".", 1), 64)
	if err != nil {
		return nil, err
	}

	return &model.DescriptionValueData{
		Description: description,
		Value:       val,
	}, nil
}
