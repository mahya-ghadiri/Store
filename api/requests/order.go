package requests

type AddToCartRequest struct {
	ProductId string `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	UserId    string `json:"user_id"`
	OrderId   string `json:"order_id"`
	SessionId int    `json:"session_id"`
}

type RemoveFromCartRequest struct {
	ProductId string `json:"product_id"`
	OrderId   string `json:"order_id"`
}
