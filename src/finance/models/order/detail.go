package order

type FinanceOrderDetail struct {
	ID       uint         `json:"-"`
	Order    FinanceOrder `gorm:"AssociationForeignKey:OrderId" json:"-"`
	OrderId  uint         `json:"-"`
	Name     string       `json:"name"`
	Quantity int          `json:"quantity"`
	Unit     int          `json:"unit"`
	Price    int          `json:"price"`
}
