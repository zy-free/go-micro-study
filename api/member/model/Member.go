package model

type Member struct {
	ID    int64  `json:"id" db:"id"`
	Phone string `json:"phone" db:"phone" `
	Name  string `json:"name" dn:"name" `
	Age   int8   `json:"age" db:"age"`
}
