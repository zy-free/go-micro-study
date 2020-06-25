package model

type ArgMemberAdd struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type ArgMemberInit struct {
	Phone string `json:"phone"  binding:"required"`
	Name  string `json:"name"  binding:"required"`
}

type ArgMemberUpdate struct {
	ID      int64  `json:"id" `
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Address string `json:"address"`
}

type ArgMemberSet struct {
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Address string `json:"address"`
}

type ArgMemberSort []struct {
	ID       int64 `json:"id"`
	OrderNum int64 `json:"order_num"`
}
