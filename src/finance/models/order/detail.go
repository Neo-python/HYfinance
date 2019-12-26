package order

type FinanceOrderDetail struct {
	ID uint `json:"-"`
	//Order    FinanceOrder `gorm:"ForeignKey:OrderId" json:"order,omitempty"`
	OrderId  uint   `json:"-"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     int    `json:"unit"`
	Price    int    `json:"price"`
}
