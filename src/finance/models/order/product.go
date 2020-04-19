package order

type Product struct {
	Name     string  `json:"name" validate:"required"`
	Quantity float64 `json:"quantity"`
	Unit     int     `json:"unit"`
	Price    float64 `json:"price"`
	Measure  float64 `json:"measure"`
}
