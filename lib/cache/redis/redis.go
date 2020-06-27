package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	xtime "github.com/zy-free/micro-study/lib/time"
)

type Config struct {
	Addr         string // redis name, for trace
	Password     string
	DialTimeout  xtime.Duration
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
	// Close items after remaining item for this duration. If the value
	// is zero, then item items are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout xtime.Duration
}

func NewPool(c *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DialTimeout:  time.Duration(c.DialTimeout),
		IdleTimeout:  time.Duration(c.IdleTimeout),
		ReadTimeout:  time.Duration(c.ReadTimeout),
		WriteTimeout: time.Duration(c.WriteTimeout),
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis addr(%s) error: %v", c.Addr, err))
	}

	return client
}
