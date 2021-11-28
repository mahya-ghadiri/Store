package handlers

import (
	"bytes"
	"context"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"store/api/requests"
	"store/internal/redis"
	"store/mocks"
	"store/models"
	"testing"
)

func TestAddToCart(t *testing.T) {

	_, redisMock := mocks.NewDataProviderMock(t)
	redis.ProductDB = redisMock
	redis.OrderDB = redisMock

	testCases := []struct {
		name               string
		request            requests.AddToCartRequest
		wantedResponseCode int
	}{
		{
			name:               "create cart and add product",
			request:            requests.AddToCartRequest{ProductId: "1", Quantity: 1, UserId: "110"},
			wantedResponseCode: http.StatusOK,
		},
	}
	productDoc := &redisearch.Document{
		Id:         "1",
		Score:      1.0,
		Properties: map[string]interface{}{"title": "book", "price": "2000"},
	}

	OrderDoc := redisearch.Document{
		Id:         "1",
		Score:      1.0,
		Properties: map[string]interface{}{"user_id": "110", "status": 1},
	}

	e := echo.New()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(mocks.JSON(testCase.request))))

			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			recorder := httptest.NewRecorder()
			c := e.NewContext(request, recorder)

			redisMock.EXPECT().GetProduct(context.Background(), "1").Return(productDoc, nil)
			redisMock.EXPECT().CreateOrder(context.Background(), models.Order{UserId: "110", Status: models.InProgress}).Return(OrderDoc, nil)
			redisMock.EXPECT().CreateOrderItem(context.Background(), models.OrderItem{ProductId: testCase.request.ProductId, OrderId: "1", Quantity: testCase.request.Quantity})

			if err := AddToCart(c); err != nil {
				t.Errorf("error happen : %v", err)
			}

			if recorder.Code != testCase.wantedResponseCode {
				t.Errorf("expecting code %v but get code %v", testCase.wantedResponseCode, recorder.Code)
			}
		})
	}
}

func TestRemoveFromCart(t *testing.T) {

	_, redisMock := mocks.NewDataProviderMock(t)
	redis.ProductDB = redisMock
	redis.OrderDB = redisMock

	testCases := []struct {
		name               string
		request            requests.RemoveFromCartRequest
		wantedResponseCode int
	}{
		{
			name:               "remove product from cart",
			request:            requests.RemoveFromCartRequest{ProductId: "1", OrderId: "2"},
			wantedResponseCode: http.StatusOK,
		},
	}
	productDoc := &redisearch.Document{
		Id:         "1",
		Score:      1.0,
		Properties: map[string]interface{}{"title": "book", "price": "2000"},
	}

	e := echo.New()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(mocks.JSON(testCase.request))))

			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			recorder := httptest.NewRecorder()
			c := e.NewContext(request, recorder)

			redisMock.EXPECT().GetProduct(context.Background(), testCase.request.ProductId).Return(productDoc, nil)
			redisMock.EXPECT().DeleteOrderItem(context.Background(), testCase.request.OrderId, testCase.request.ProductId).Return(nil)

			if err := RemoveFromCart(c); err != nil {
				t.Errorf("error happen : %v", err)
			}

			if recorder.Code != testCase.wantedResponseCode {
				t.Errorf("expecting code %v but get code %v", testCase.wantedResponseCode, recorder.Code)
			}
		})
	}
}
