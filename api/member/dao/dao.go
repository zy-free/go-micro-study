package dao

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/zy-free/micro-study/lib/orm"
)

type Dao struct {
	member     *gorm.DB
	memberRead *gorm.DB
	redis      *redis.Pool
	// memcache       *memcache.Pool
	// es             *elastic.Elastic
	// hbase
	// tidb
}

func New() *Dao {

	d := &Dao{
		member: orm.NewMySQL(&orm.Config{DSN: "username:password@tcp(host)/db_name?charset=utf8&allowOldPasswords=1"}),
	}
	return d
}
