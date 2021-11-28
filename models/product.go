package models

import (
	"context"
	"github.com/RediSearch/redisearch-go/redisearch"
)

type Product struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type ProductProvider interface {
	CreateProductIndex(ctx context.Context) error
	CreateProduct(ctx context.Context, product Product) (redisearch.Document, error)
	SearchProductByTitle(ctx context.Context, title string) ([]redisearch.Document, error)
	GetProduct(ctx context.Context, id string) (*redisearch.Document, error)
}
