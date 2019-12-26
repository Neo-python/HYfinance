package receiver

import (
	"time"
)

type FinanceReceiver struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	Tel       string     `json:"tel"`
}

type FinanceReceiverJson struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}
