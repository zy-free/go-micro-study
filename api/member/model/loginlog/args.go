package loginlog

type ArgLogAdd struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
