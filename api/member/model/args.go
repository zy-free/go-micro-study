package model

type ArgMemberAdd struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type ArgMemberGet struct {
	ID int64 `form:"id" binding:"required,gt=0"`
}

type ArgMemberInit struct {
	Phone string `json:"phone"  binding:"required"`
	Name  string `json:"name"  binding:"required"`
}

type ArgMemberUpdate struct {
	ID      int64  `json:"id" binding:"required,gt=0"`
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Address string `json:"address"`
}

type ArgMemberSet struct {
	ID      int64  `json:"id" binding:"required,gt=0"`
	Phone   string `json:"phone" binding:"required,max=20"`
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Address string `json:"address"`
}

type ArgMemberSort []struct {
	ID       int64 `json:"id" binding:"required,gt=0"`
	OrderNum int64 `json:"order_num" binding:"required,gt=0"`
}

type ArgMemberDelete struct {
	ID int64 `form:"id" binding:"required,gt=0"`
}
