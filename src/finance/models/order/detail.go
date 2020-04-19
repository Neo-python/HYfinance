package order

type FinanceOrderDetail struct {
	ID       uint         `json:"-"`
	Order    FinanceOrder `gorm:"AssociationForeignKey:OrderId" json:"-"`
	OrderId  uint         `json:"-" gorm:"COMMENT:'订单编号'"`
	Name     string       `json:"name" gorm:"COMMENT:'产品名'"`
	Quantity float64      `json:"quantity" gorm:"COMMENT:'数量'"`
	Unit     int          `json:"unit" gorm:"COMMENT:'单位'"`
	Price    float64      `json:"price" gorm:"COMMENT:'价格'"`
	Measure  float64      `json:"measure" gorm:"COMMENT:'计量单位值,测量值'"`
}
