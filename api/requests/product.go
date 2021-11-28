package requests


type CreateProductRequest struct {
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
}

