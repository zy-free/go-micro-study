package dao

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/zy-free/micro-study/api/member/conf"
	"github.com/zy-free/micro-study/lib/database/mysql"
)

type Dao struct {
	//member     *gorm.DB
	//memberRead *gormDB
	member *sqlx.DB
	memberRead *sqlx.DB

	redis      *redis.Pool
	// memcache       *memcache.Pool
	// es             *elastic.Elastic
	// hbase
	// tidb
}

func New(c *conf.Config) *Dao {
	d := &Dao{
		member: mysql.NewMySQL(c.Mysql.Member),
		memberRead: mysql.NewMySQL(c.Mysql.MemberRead),
	}
	return d
}
