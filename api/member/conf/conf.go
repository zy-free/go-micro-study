package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"github.com/zy-free/micro-study/lib/cache/redis"
	"github.com/zy-free/micro-study/lib/database/mysql"
	xtime "github.com/zy-free/micro-study/lib/time"
	"os"
)

var (
	confPath string
	Conf     = &Config{}
)

type Config struct {
	// db
	Mysql *Mysql
	Redis *Redis
}

// Mysql is mysql conf
type Mysql struct {
	Member     *mysql.Config
	MemberRead *mysql.Config
}

type Redis struct {
	*redis.Config
	Expire xtime.Duration
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init conf
func Init() (err error) {
	if confPath != "" {
		return local() // golang 相对路径为当前命令执行的相对路径
	}
	return remote()
}

func local() (err error) {
	//spew.Dump("local confPath:",confPath)
	_, err = toml.DecodeFile(confPath, &Conf)
	//spew.Dump("local config init:",Conf)
	return
}

func remote() (err error) {
	confPath = os.Getenv("CONFIG_PATH") // micro-study/api/member/conf/test.toml
	spew.Dump("remote confPath:", confPath)
	_, err = toml.DecodeFile(confPath, &Conf)
	spew.Dump("remote config init:", Conf)
	// todo 引入配置中心 实际环境读取不同配置
	return nil
}
