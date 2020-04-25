package handler

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestAPI_Crawler(t *testing.T) {
	type TestMock struct {
		name         string
		path         string
		url          string
		expectedCode int
		expectedBody string
	}
	tt := []TestMock{
		{
			name:         "wrong path request",
			expectedCode: http.StatusNotFound,
			expectedBody: "404 page not found\n",
		},
		{
			name:         "wrong url",
			expectedCode: http.StatusInternalServerError,
			path:         "/smartmei/crawler",
			expectedBody: "{\"result\":false,\"code\":-1,\"message\":\"Unknown error: result: false; code: 4010009; message: query params `url` is not valid\"}\n",
		},
		{
			name:         "error getting info",
			expectedCode: http.StatusInternalServerError,
			url:          "https://www.google.com",
			path:         "/smartmei/crawler",
			expectedBody: "{\"result\":false,\"code\":-1,\"message\":\"Unknown error: result: false; code: 4010005; message: error crawling data: result: false; code: 4010008; message: found array size 1, but array size 2 is required\"}\n",
		},
		{
			name:         "all ok",
			expectedCode: http.StatusOK,
			url:          "https://www.smartmei.com.br",
			path:         "/smartmei/crawler",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			api := New()
			exp := httpexpect.WithConfig(httpexpect.Config{
				Client: &http.Client{
					Transport: httpexpect.NewBinder(api),
					Jar:       httpexpect.NewJar(),
				},
				Reporter: httpexpect.NewAssertReporter(t),
			})

			if tc.expectedCode != http.StatusOK {
				exp.GET(tc.path).
					WithQuery("url", tc.url).
					Expect().
					Status(tc.expectedCode).
					Body().Equal(tc.expectedBody)
			} else {
				// just checking if it is not empty, since I don't know exactly what to expect
				// as response body
				exp.GET(tc.path).
					WithQuery("url", tc.url).
					Expect().
					Status(tc.expectedCode).
					Body().NotEmpty()
			}
		})
	}
}
