package dao

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zy-free/micro-study/api/member/model"
)

const (
	_memberKey = "m:%d" // set r:mid
)

// todo 参考favorite 然后依靠databus设置缓存，那么micro就要自己照消息组件

func memberKey(id int64) string {
	return fmt.Sprintf(_memberKey, id)
}

func (dao *Dao) ExpireMemberCache(id int64) (ok bool, err error) {
	key := memberKey(id)
	if ok, err = dao.redis.Expire(key, dao.redisExpire).Result(); err != nil {
		return false, errors.Wrapf(err, "ExpireMember id(%d)", id)
	}
	return
}

func (dao *Dao) GetMemberCache(id int64) (member *model.Member, err error) {
	key := memberKey(id)
	member = &model.Member{}
	if err = dao.redis.Get(key).Scan(member); err != nil {
		return nil, errors.Wrapf(err, "GetMemberCache id(%d)", id)
	}
	return
}

func (dao *Dao) AddMemberCache(id int64, member *model.Member) (err error) {
	key := memberKey(id)
	if err = dao.redis.Set(key, member, dao.redisExpire).Err(); err != nil {
		return errors.Wrapf(err, "AddMemberCache id(%d),member(%v)", id, member)
	}
	return
}

func (dao *Dao) DelMemberCache(id int64) (err error) {
	key := memberKey(id)
	if err = dao.redis.Del(key).Err(); err != nil {
		return errors.Wrapf(err, "DelMemberCache id(%d)", id)
	}
	return
}
