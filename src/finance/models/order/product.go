package order

type Product struct {
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity"`
	Unit     int    `json:"unit"`
	Price    int    `json:"price"`
	Measure  int    `json:"measure"`
}
