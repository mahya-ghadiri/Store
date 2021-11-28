package handlers

import (
	"bytes"
	"context"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"net/url"
	"store/api/requests"
	"store/internal/redis"
	"store/mocks"
	"store/models"
	"testing"
)

func TestCreateProduct(t *testing.T) {

	_, redisMock := mocks.NewDataProviderMock(t)
	redis.ProductDB = redisMock

	testCases := []struct {
		name               string
		body               string
		wantedResponseCode int
	}{
		{
			name:               "wrong json format error happens from bind",
			body:               `{"title":"black friday}`,
			wantedResponseCode: http.StatusBadRequest,
		},
		{
			name:               "successful creation",
			body:               mocks.JSON(requests.CreateProductRequest{Title: "book", Price: 2000}),
			wantedResponseCode: http.StatusOK,
		},
	}
	productDoc := redisearch.Document{
		Id:         "1",
		Score:      1.0,
		Properties: map[string]interface{}{"title": "book", "price": "2000"}}

	product := models.Product{
		Title: "book",
		Price: 2000,
	}

	redisMock.EXPECT().CreateProduct(context.Background(), product).Return(productDoc, nil)

	e := echo.New()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(testCase.body)))

			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			recorder := httptest.NewRecorder()
			c := e.NewContext(request, recorder)

			if err := CreateProduct(c); err != nil {
				t.Errorf("error happen : %v", err)
			}

			if recorder.Code != testCase.wantedResponseCode {
				t.Errorf("expecting code %v but get code %v", testCase.wantedResponseCode, recorder.Code)
			}
		})
	}
}

func TestSearchProduct(t *testing.T) {
	_, redisMock := mocks.NewDataProviderMock(t)
	redis.ProductDB = redisMock

	testCases := []struct {
		name               string
		query              string
		wantedResponseCode int
		result             []redisearch.Document
	}{
		{
			name:               "product not exists with this title",
			query:              "sss",
			wantedResponseCode: http.StatusNotFound,
			result:             []redisearch.Document{},
		},
		{
			name:               "successful search",
			query:              "book",
			wantedResponseCode: http.StatusOK,
			result:             []redisearch.Document{{"1", 1.0, nil, map[string]interface{}{"title": "book", "price": "2000"}}},
		},
	}

	e := echo.New()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			redisMock.EXPECT().SearchProductByTitle(context.Background(), testCase.query).Return(testCase.result, nil)

			q := make(url.Values)
			q.Set("query", testCase.query)
			request := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)

			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			recorder := httptest.NewRecorder()
			c := e.NewContext(request, recorder)

			if err := SearchProduct(c); err != nil {
				t.Errorf("error happen : %v", err)
			}

			if recorder.Code != testCase.wantedResponseCode {
				t.Errorf("expecting code %v but get code %v", testCase.wantedResponseCode, recorder.Code)
			}
		})
	}
}
