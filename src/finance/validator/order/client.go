package order

type QueryForm struct {
	Name  string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
}
