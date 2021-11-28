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

func (redisClient *Client) CreateProductIndex(ctx context.Context) error {
	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true}))

	if err := redisClient.db.Drop(); err != nil {
		fmt.Println(" redisClient.db.Drop()", err)
	}

	if err := redisClient.db.CreateIndex(schema); err != nil {
		fmt.Println("redisClient.db.CreateIndex", err)
		return err
	}

	return nil
}

func (redisClient *Client) CreateProduct(ctx context.Context, product models.Product) (redisearch.Document, error) {
	id := rand.Intn(100000)
	doc := redisearch.NewDocument(strconv.Itoa(id), 1.0).
		Set("title", product.Title).
		Set("price", product.Price).
		Set("created_at", time.Now().Unix()).
		Set("updated_at", time.Now().Unix())

	if err := redisClient.db.Index([]redisearch.Document{doc}...); err != nil {
		fmt.Println("Index", err)
		return redisearch.Document{}, err
	}
	return doc, nil
}

func (redisClient *Client) SearchProductByTitle(ctx context.Context, title string) ([]redisearch.Document, error) {
	docs, total, err := redisClient.db.Search(redisearch.NewQuery(title).
		Limit(0, 2).
		SetReturnFields("title", "price"))
	if err != nil {
		return nil, err
	}
	fmt.Println(docs, total, err)
	if len(docs) > 0 {
		fmt.Println(docs[0].Id, docs[0].Properties["title"], total, err)
		fmt.Println(docs)
	}
	return docs, err
}

func (redisClient *Client) GetProduct(ctx context.Context, id string) (*redisearch.Document, error) {
	doc, err := redisClient.db.Get(id)
	if err != nil {
		return nil, err
	}

	return doc, err
}
