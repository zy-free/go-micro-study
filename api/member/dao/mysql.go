package dao

import "github.com/zy-free/micro-study/api/member/model"

func (dao *Dao) CreateMember(arg *model.ArgMemberCreate) (int64, error) {
	if err := dao.member.Table("member").Create(arg).Error; err != nil {
		return 0, err
	}
	return int64(arg.ID), nil
}
