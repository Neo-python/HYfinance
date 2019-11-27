package order

type OrderDetail struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     int    `json:"unit"`
	Price    int    `json:"price"`
}
