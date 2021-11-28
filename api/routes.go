package api

import (
	"github.com/labstack/echo/v4"
	"store/api/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", handlers.HealthCheck)

	// store
	e.POST("/product", handlers.CreateProduct)
	e.GET("/product/search", handlers.SearchProduct)
	e.POST("/cart/add", handlers.AddToCart)
	e.POST("/cart/remove", handlers.RemoveFromCart)

}
