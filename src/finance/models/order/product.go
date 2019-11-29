package order

type Product struct {
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
