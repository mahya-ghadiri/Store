package handlers

import (
	"context"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/labstack/echo/v4"
	"net/http"
	"store/api/requests"
	"store/api/responses"
	"store/internal/redis"
	"store/models"
)

func AddToCart(ctx echo.Context) (err error) {
	request := new(requests.AddToCartRequest)

	if err = ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewError(http.StatusBadRequest, err.Error()))
	}

	product, err := redis.ProductDB.GetProduct(context.Background(), request.ProductId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewError(http.StatusInternalServerError, err.Error()))
	}

	if product == nil {
		return ctx.JSON(http.StatusNotFound, responses.NewError(http.StatusNotFound, "This product does not exists"))
	}

	var orderDocument *redisearch.Document
	if request.OrderId != "" {
		orderDocument, err = redis.OrderDB.GetOrder(context.Background(), request.OrderId)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, responses.NewError(http.StatusInternalServerError, err.Error()))
		}
	}

	orderId := request.OrderId
	// if cart not exists create cart
	if orderDocument == nil {
		order := models.Order{
			UserId:    request.UserId,
			Status:    models.InProgress,
		}
		createdOrder, err := redis.OrderDB.CreateOrder(context.Background(), order)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, responses.NewError(http.StatusInternalServerError, err.Error()))
		}
		orderId = createdOrder.Id
	}

	orderItem := models.OrderItem{
		OrderId:   orderId,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	}
	err = redis.OrderDB.CreateOrderItem(context.Background(), orderItem)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewError(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, responses.NewDefault(http.StatusOK, "Product Added to cart Successfully", map[string]string{"id": orderId}))
}

func RemoveFromCart(ctx echo.Context) (err error) {

	request := new(requests.AddToCartRequest)

	if err = ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewError(http.StatusBadRequest, err.Error()))
	}

	product, err := redis.ProductDB.GetProduct(context.Background(), request.ProductId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewError(http.StatusInternalServerError, err.Error()))
	}

	if product == nil {
		return ctx.JSON(http.StatusNotFound, responses.NewError(http.StatusNotFound, "This product does not exists"))
	}

	err = redis.OrderDB.DeleteOrderItem(context.Background(), request.OrderId, request.ProductId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewError(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, responses.NewDefault(http.StatusOK, "Product removed from cart Successfully", nil))
}
