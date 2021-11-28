package redis

import (
	"store/models"
)

type DataProvider interface {
	models.ProductProvider
	models.OrderProvider
	models.OrderItemProvider
}
