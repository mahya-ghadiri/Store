package models

import (
	"context"
	"github.com/RediSearch/redisearch-go/redisearch"
)

type Order struct {
	Id        string      `json:"id"`
	UserId    string      `json:"user_id"`
	SessionId int         `json:"session_id"`
	Status    OrderStatus `json:"status"`
}

type OrderProvider interface {
	CreateOrderIndex(ctx context.Context) error
	CreateOrder(ctx context.Context, order Order) (redisearch.Document, error)
	GetOrder(ctx context.Context, id string) (*redisearch.Document, error)
}

type OrderStatus uint8

const (
	InProgress OrderStatus = iota + 1
	Paid
	Sent
)
