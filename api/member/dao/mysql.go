package dao

import (
	"bytes"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/zy-free/micro-study/api/member/model"
	"github.com/zy-free/micro-study/lib/xstr"
	"strconv"
	"strings"
	"time"
)

func (dao *Dao) InitMember(arg *model.ArgMemberInit) (err error) {
	now := time.Now()
	sql := `INSERT IGNORE INTO member (phone,name,created_at,updated_at) VALUES(?,?,?,?)`
	if _, err = dao.member.Exec(sql, arg.Phone, arg.Name, now, now); err != nil {
		return errors.Wrapf(err, "InitMember arg(%v)", arg)
	}
	return
}

// 创建单个
func (dao *Dao) AddMember(arg *model.ArgMemberAdd) (id int64, err error) {
	now := time.Now()
	sql := `INSERT INTO member (phone,name,created_at,updated_at) VALUES(?,?,?,?)`
	result, err := dao.member.Exec(sql, arg.Phone, arg.Name, now, now)
	if err != nil {
		return 0, errors.Wrapf(err, "AddMember arg(%v)", arg)
	}

	id, _ = result.LastInsertId()
	return
}

// 批量创建
func (dao *Dao) BatchAddMember(args []*model.ArgMemberAdd) (affectRow int64, err error) {
	now := time.Now()
	sql := `INSERT INTO member (phone,name,created_at,updated_at) VALUES `
	var valueString []string
	var valueArgs []interface{}
	for _, arg := range args {
		valueString = append(valueString, "(?,?,?,?)")
		valueArgs = append(valueArgs, arg.Phone, arg.Name, now, now)
	}
	result, err := dao.member.Exec(sql+strings.Join(valueString, ","), valueArgs...)
	if err != nil {
		return 0, errors.Wrapf(err, "BatchAddMember arg(%v)", args)
	}
	affectRow, _ = result.RowsAffected()
	return
}

// 删除
func (dao *Dao) DeleteMember(id int64) (err error) {
	now := time.Now()
	sql := `UPDATE member SET deleted = 1 , deleted_at = ? WHERE id = ?`
	if _, err = dao.member.Exec(sql, now, id); err != nil {
		return errors.Wrapf(err, "DeleteMember id(%d)", id)
	}
	return
}

// 更新单个
func (dao *Dao) UpdateMember(arg *model.ArgMemberUpdate) (err error) {
	sql := "UPDATE member SET name =:name, "
	updateMap := map[string]interface{}{}
	updateMap["id"] = arg.ID
	updateMap["name"] = arg.Name
	if arg.Address != "" {
		sql += " address =:address,"
		updateMap["address"] = arg.Address
	}
	if arg.Age != 0 {
		sql += " age =:age"
		updateMap["age"] = arg.Age
	}
	sql += " WHERE id =:id"
	if _, err = dao.member.NamedExec(sql, updateMap); err != nil {
		return errors.Wrapf(err, "UpdateMember arg(%v)", arg)
	}
	return
}

// 更新或创建
func (dao *Dao) SetMember(arg *model.ArgMemberSet) (err error) {
	now := time.Now()
	sql := "INSERT INTO member (id,phone,name,age,address,created_at,updated_at) VALUES (?,?,?,?,?,?,?) " +
		"ON DUPLICATE KEY UPDATE phone=?,name=?,age=?,address=?,updated_at=?"
	if _, err = dao.member.Exec(sql, arg.ID, arg.Phone, arg.Name, arg.Age, arg.Address, now, now, arg.Phone, arg.Name, arg.Age, arg.Address, now); err != nil {
		return errors.Wrapf(err, "SetMember arg(%v)", arg)
	}
	return
}

// 批量更改顺序
func (dao *Dao) BatchUpdateMemberSort(args model.ArgMemberSort) (err error) {
	var (
		buf bytes.Buffer
		ids []int64
	)
	buf.WriteString("UPDATE member SET order_num = CASE id")
	for _, arg := range args {
		buf.WriteString(" WHEN ")
		buf.WriteString(strconv.FormatInt(arg.ID, 10))
		buf.WriteString(" THEN ")
		buf.WriteString(strconv.FormatInt(arg.OrderNum, 10))
		ids = append(ids, arg.ID)
	}
	buf.WriteString(" END  WHERE id IN (")
	buf.WriteString(xstr.JoinInts(ids))
	buf.WriteString(")")
	if _, err = dao.member.Exec(buf.String()); err != nil {
		return errors.Wrapf(err, "BatchUpdateMemberOrder args(%v)", args)
	}
	return
}

// 根据id查询单个
func (dao *Dao) GetMemberByID(id int64) (m *model.Member, err error) {
	m = &model.Member{}
	sql := `SELECT id,phone,name FROM member WHERE id = ? AND deleted = 0 `
	if err = dao.member.Get(m, sql, id); err != nil {
		return nil, errors.Wrapf(err, "GetMemberByID id(%d)", id)
	}
	return
}

func (dao *Dao) GetMemberMaxAge() (age int8, err error) {
	sql := `SELECT IFNULL(MAX(age),0) FROM member WHERE deleted = 0 `
	if err = dao.member.Get(&age, sql); err != nil {
		return 0, errors.Wrapf(err, "GetMemberMaxAge")
	}
	return
}

func (dao *Dao) GetMemberSumAge() (age int64, err error) {
	sql := `SELECT IFNULL(SUM(age),0) FROM member WHERE deleted = 0`
	if err = dao.member.Get(&age, sql); err != nil {
		return 0, errors.Wrapf(err, "GetMemberMaxAge")
	}
	return
}

func (dao *Dao) CountMember() (count int64, err error) {
	sql := `SELECT COUNT(*) FROM member `
	if err = dao.member.Get(&count, sql); err != nil {
		return 0, errors.Wrapf(err, "CountMember")
	}
	return
}

func (dao *Dao) HasMemberByID(id int64) (has bool, err error) {
	var count int
	s := `SELECT COUNT(*) FROM member WHERE id = ? `
	err = dao.member.Get(&count, s, id)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrapf(err, "HasMemberByID id(%d)", id)
	}
	return count > 0, nil
}

// 根据其他属性查询列表
func (dao *Dao) QueryMemberByName(name string) (res []*model.Member, err error) {
	res = make([]*model.Member, 0, 0) // 返回nil还是空切片会影响json里的结构
	sql := `SELECT phone,name FROM member WHERE name = ? AND deleted = 0 `
	if err = dao.member.Select(&res, sql, name); err != nil {
		return res, errors.Wrapf(err, "QueryMemberByName name(%s)", name)
	}
	return
}

// 根据ids查询列表
func (dao *Dao) QueryMemberByIDs(ids []int64) (res map[int64]*model.Member, err error) {
	var t []*model.Member
	res = make(map[int64]*model.Member)
	sql := "SELECT id,phone,name FROM member WHERE id IN (" + xstr.JoinInts(ids) + ") AND deleted = 0 "
	if err = dao.member.Select(&t, sql); err != nil {
		return res, errors.Wrapf(err, "QueryMemberByIDs ids(%v)", ids)
	}
	for _, r := range t {
		res[r.ID] = r
	}
	return
}
