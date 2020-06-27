package orm

import (
	"fmt"
	"time"

	// database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	xtime "github.com/zy-free/micro-study/lib/time"
)

// Config mysql config.
type Config struct {
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout xtime.Duration // connect max life time.
}

// 自定义error
//func init() {
//	gorm.ErrRecordNotFound = ecode.NothingFound
//}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		panic(fmt.Sprintf("db dsn(%s) error: %v", c.DSN, err))
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	return
}
