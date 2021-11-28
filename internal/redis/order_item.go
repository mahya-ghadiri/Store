package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"math/rand"
	"store/models"
	"strconv"
	"time"
)

func (redisClient *Client) CreateOrderItemIndex(ctx context.Context) error {
	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("order_id", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewTextFieldOptions("product_id", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true}))

	if err := redisClient.db.Drop(); err != nil {
		fmt.Println(" redisClient.db.Drop()", err)
	}

	if err := redisClient.db.CreateIndex(schema); err != nil {
		fmt.Println("redisClient.db.CreateIndex", err)
		return err
	}

	return nil
}

func (redisClient *Client) CreateOrderItem(ctx context.Context, item models.OrderItem) error {
	id := rand.Intn(100000)
	doc := redisearch.NewDocument(strconv.Itoa(id), 1.0).
		Set("order_id", item.OrderId).
		Set("product_id", item.ProductId).
		Set("quantity", item.Quantity).
		Set("created_at", time.Now().Unix()).
		Set("updated_at", time.Now().Unix())

	if err := redisClient.db.Index([]redisearch.Document{doc}...); err != nil {
		return err
	}
	return nil
}

func (redisClient *Client) DeleteOrderItem(ctx context.Context, orderId string, productId string) error {
	query := fmt.Sprintf("@order_id:%s @product_id:%s", orderId, productId)
	docs, _, err := redisClient.db.Search(redisearch.NewQuery(query))
	if err != nil {
		return err
	}

	if len(docs) == 0 {
		return errors.New("order not found")
	}

	if err := redisClient.db.DeleteDocument(docs[0].Id); err != nil {
		return err
	}
	return nil
}
