package service

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/dexterorion/smartmei/model"

	"github.com/dexterorion/smartmei/errors"

	"github.com/stretchr/testify/require"
)

func TestService_GetConvertedData(t *testing.T) {
	type TestExecution struct {
		name          string
		expectedError error
		url           string
		from          string
		to            []string
	}

	tt := []TestExecution{
		{
			name:          "build request error",
			url:           "---",
			expectedError: errors.ErrorCrawlingData(errors.ErrorDoingRequest(fmt.Errorf(`Get ---: unsupported protocol scheme ""`))),
		},
		{
			name:          "invalid status code",
			url:           "https://google.com",
			expectedError: errors.ErrorCrawlingData(errors.InvalidArraySize(2, 1)),
		},
		{
			name:          "all ok",
			url:           "https://www.smartmei.com.br",
			from:          "BLABLA",
			to:            []string{"ABC", "DEF"},
			expectedError: errors.ErrorGettingExchange(errors.InvalidStatusCode(http.StatusOK, http.StatusBadRequest)),
		},
		{
			name: "all ok",
			url:  "https://www.smartmei.com.br",
			from: model.CurrencyBRL,
			to:   []string{model.CurrencyUSD, model.CurrencyEUR},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sv := New()

			result, err := sv.GetConvertedData(context.Background(), tc.url, tc.from, tc.to)
			if err != nil {
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NotEqual(t, 0, result.Rates.Eur)
				require.NotEqual(t, 0, result.Rates.Brl)
				require.NotEqual(t, 0, result.Rates.Usd)
			}
		})
	}
}
