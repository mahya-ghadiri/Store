package models

import "context"

type OrderItem struct {
	Id        string `json:"id"`
	OrderId   string `json:"order_id"`
	ProductId string `json:"product_id"`
	Quantity  uint   `json:"quantity"`
}

type OrderItemProvider interface {
	CreateOrderItemIndex(ctx context.Context) error
	CreateOrderItem(ctx context.Context, item OrderItem) error
	DeleteOrderItem(ctx context.Context, orderId string, productId string) error
}
