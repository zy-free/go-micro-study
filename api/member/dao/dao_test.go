package dao

import (
	"flag"
	"github.com/zy-free/micro-study/api/member/conf"
	"os"
	"path/filepath"
	"testing"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	dir ,_ :=  filepath.Abs("../conf/test.toml")
	flag.Set("conf",dir)
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	m.Run()
	os.Exit(0)
}
