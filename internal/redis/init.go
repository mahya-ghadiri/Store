package redis

import (
	"context"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"store/internal/config"
)

type Client struct {
	db *redisearch.Client
}

// implementation of DataProvider interface.
var ProductDB DataProvider
var OrderDB DataProvider

func Connect(redisConfig config.Redis) (*redisearch.Client, *redisearch.Client) {
	productClient := redisearch.NewClient(redisConfig.Address, "productIndex")
	orderClient := redisearch.NewClient(redisConfig.Address, "orderIndex")
	return productClient, orderClient
}

func Init(productClient *redisearch.Client, orderClient *redisearch.Client) {
	ProductDB = &Client{db: productClient}
	OrderDB = &Client{db: orderClient}
	err := ProductDB.CreateProductIndex(context.Background())
	fmt.Println(err)
	err = OrderDB.CreateOrderItemIndex(context.Background())
	fmt.Println(err)
}
