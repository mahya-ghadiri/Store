package redis

import (
	"context"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"math/rand"
	"store/models"
	"strconv"
	"time"
)

func (redisClient *Client) CreateOrderIndex(ctx context.Context) error {
	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("user_id", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("session_id")).
		AddField(redisearch.NewNumericField("status")).
		AddField(redisearch.NewNumericField("created_at")).
		AddField(redisearch.NewNumericField("updated_at"))

	if err := redisClient.db.Drop(); err != nil {
		fmt.Println(" redisClient.db.Drop()", err)
	}

	if err := redisClient.db.CreateIndex(schema); err != nil {
		fmt.Println("redisClient.db.CreateIndex", err)
		return err
	}

	return nil
}

func (redisClient *Client) CreateOrder(ctx context.Context, order models.Order) (redisearch.Document, error) {
	id := rand.Intn(100000)
	doc := redisearch.NewDocument(strconv.Itoa(id), 1.0).
		Set("user_id", order.UserId).
		Set("status", order.Status).
		Set("created_at", time.Now().Unix()).
		Set("updated_at", time.Now().Unix())

	if err := redisClient.db.Index([]redisearch.Document{doc}...); err != nil {
		return redisearch.Document{}, err
	}
	return doc, nil
}

func (redisClient *Client) GetOrder(ctx context.Context, id string) (*redisearch.Document, error) {
	doc, err := redisClient.db.Get(id)
	if err != nil {
		return nil, err
	}

	return doc, err
}
