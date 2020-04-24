package errors

import (
	"fmt"

	"github.com/dexterorion/smartmei/lib"
)

const (
	errorPrefix                 = 4010
	errorBuildingRequest        = 1
	errorDoingRequest           = 2
	errorProcessingResponseBody = 3
	invalidStatusCode           = 4
	errorCrawlingData           = 5
	errorGettingExchange        = 6
	errorConvertingByteToMap    = 7
	invalidArraySize            = 8
	urlCannotBeEmpty            = 9
)

// ErrorBuildingRequest represents an error when building request
func ErrorBuildingRequest(err error) error {
	return lib.NewError(errorPrefix, errorBuildingRequest, fmt.Errorf("error building request: %s", err.Error()))
}

// ErrorDoingRequest represents an error when doing request
func ErrorDoingRequest(err error) error {
	return lib.NewError(errorPrefix, errorDoingRequest, fmt.Errorf("error doing request: %s", err.Error()))
}

// ErrorProcessingResponseBody represents an error when doing request
func ErrorProcessingResponseBody(err error) error {
	return lib.NewError(errorPrefix, errorProcessingResponseBody, fmt.Errorf("error processing response body: %s", err.Error()))
}

// InvalidStatusCode represents an error when unexpected status code happens
func InvalidStatusCode(expected, found int) error {
	return lib.NewError(errorPrefix, invalidStatusCode, fmt.Errorf("found status code %d, but status code %d is required", found, expected))
}

// ErrorCrawlingData represents an error when crawling data
func ErrorCrawlingData(err error) error {
	return lib.NewError(errorPrefix, errorCrawlingData, fmt.Errorf("error crawling data: %s", err.Error()))
}

// ErrorGettingExchange represents an error when getting exchange data
func ErrorGettingExchange(err error) error {
	return lib.NewError(errorPrefix, errorGettingExchange, fmt.Errorf("error getting exchange data: %s", err.Error()))
}

// ErrorConvertingByteToMap represents an error when converting bytes to map
func ErrorConvertingByteToMap(err error) error {
	return lib.NewError(errorPrefix, errorConvertingByteToMap, fmt.Errorf("error converting byte to map: %s", err.Error()))
}

// InvalidArraySize represents an error when unexpected array size happens
func InvalidArraySize(expected, found int) error {
	return lib.NewError(errorPrefix, invalidArraySize, fmt.Errorf("found array size %d, but array size %d is required", found, expected))
}

// URLCannotBeEmpty represents an error when the url on query request is empty
func URLCannotBeEmpty() error {
	return lib.NewError(errorPrefix, urlCannotBeEmpty, fmt.Errorf("query params `url` is not valid"))
}
