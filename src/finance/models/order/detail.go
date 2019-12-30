package order

type FinanceOrderDetail struct {
	ID       uint         `json:"-"`
	Order    FinanceOrder `gorm:"AssociationForeignKey:OrderId" json:"-"`
	OrderId  uint         `json:"-" gorm:"COMMENT:'订单编号'"`
	Name     string       `json:"name" gorm:"COMMENT:'产品名'"`
	Quantity int          `json:"quantity" gorm:"COMMENT:'数量'"`
	Unit     int          `json:"unit" gorm:"COMMENT:'单位'"`
	Price    int          `json:"price" gorm:"COMMENT:'价格'"`
}
