package loginlog

import (
	"github.com/zy-free/micro-study/api/member/conf"
)

type Dao struct {
	//member     *gorm.DB
	//memberRead *gormDB
	//member     *sqlx.DB
	//memberRead *sqlx.DB
	//
	//redis       *redis.Client
	//redisExpire time.Duration

	// memcache       *memcache.Pool
	// es             *elastic.Elastic
	// hbase
	// tidb
}

func New(c *conf.Config) *Dao {
	d := &Dao{}
	return d
}
