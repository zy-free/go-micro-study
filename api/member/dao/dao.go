package dao

import (
	redis "github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/zy-free/micro-study/api/member/conf"
	xredis "github.com/zy-free/micro-study/lib/cache/redis"
	"github.com/zy-free/micro-study/lib/database/mysql"
	"time"
)

type Dao struct {
	//member     *gorm.DB
	//memberRead *gormDB
	member     *sqlx.DB
	memberRead *sqlx.DB

	redis       *redis.Client
	redisExpire time.Duration

	// memcache       *memcache.Pool
	// es             *elastic.Elastic
	// hbase
	// tidb
}

func New(c *conf.Config) *Dao {
	d := &Dao{
		// mysql
		member:     mysql.NewMySQL(c.Mysql.Member),
		memberRead: mysql.NewMySQL(c.Mysql.MemberRead),
		// redis
		redis:       xredis.NewPool(c.Redis.Config),
		redisExpire: time.Duration(c.Redis.Expire),
	}
	return d
}
