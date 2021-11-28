package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"store/api/requests"
	"store/api/responses"
	"store/internal/redis"
	"store/models"
)

func CreateProduct(ctx echo.Context) (err error) {
	request := new(requests.CreateProductRequest)

	if err = ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewError(http.StatusBadRequest, err.Error()))
	}

	product := models.Product{
		Title: request.Title,
		Price: request.Price,
	}

	productDoc, err := redis.ProductDB.CreateProduct(context.Background(), product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewError(http.StatusInternalServerError, err.Error()))

	}
	return ctx.JSON(http.StatusOK, responses.NewDefault(http.StatusOK, "Product Created Successfully", map[string]string{"id": productDoc.Id}))
}

func SearchProduct(ctx echo.Context) (err error) {
	query := ctx.QueryParam("query")

	docs, err := redis.ProductDB.SearchProductByTitle(context.Background(), query)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewError(http.StatusInternalServerError, err.Error()))
	}

	if len(docs) == 0 {
		return ctx.JSON(http.StatusNotFound, responses.NewError(http.StatusNotFound, "Product with this title not found"))
	}

	var products []responses.SearchProductResponse
	for _, doc := range docs {
		product := responses.SearchProductResponse{
			Title: doc.Properties["title"].(string),
			Price: doc.Properties["price"].(string),
		}
		products = append(products, product)
	}

	return ctx.JSON(http.StatusOK, responses.NewDefault(http.StatusOK, "", map[string]interface{}{"products": products}))
}
