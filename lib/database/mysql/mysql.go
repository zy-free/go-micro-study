package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	xtime "github.com/zy-free/micro-study/lib/time"
	"time"

	// database driver
	_ "github.com/go-sql-driver/mysql"
)

// Config mysql config.
type Config struct {
	DSN          string          // write data source name.
	Active       int             // pool
	Idle         int             // pool
	IdleTimeout  xtime.Duration   // connect max life time.
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *sqlx.DB) {
	db, err := sqlx.Open("mysql", c.DSN)
	if err != nil {
		panic(fmt.Sprintf("db dsn(%s) error: %v", c.DSN, err))
	}
	db.SetMaxIdleConns(c.Idle)
	db.SetMaxOpenConns(c.Active)
	db.SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	return
}