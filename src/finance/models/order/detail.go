package order

type FinanceOrderDetail struct {
	ID       uint
	Order    FinanceOrder `gorm:"ForeignKey:OrderId"`
	OrderId  uint         `json:"order_id"`
	Name     string       `json:"name"`
	Quantity int          `json:"quantity"`
	Unit     int          `json:"unit"`
	Price    int          `json:"price"`
}
