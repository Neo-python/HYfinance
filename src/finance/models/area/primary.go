package area

type Area struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	SuperiorId uint   `json:"superior_id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
}
