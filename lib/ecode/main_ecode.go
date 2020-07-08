package ecode

// main ecode interval is [0,990000]
var (
	// member
	MemberAdd    = New(10001, "用户添加失败")
	MemberGet    = New(10002, "用户查询失败")
	MemberUpdate = New(10003, "用户更新失败")
	MemberDelete = New(10004, "用户删除失败")
)
