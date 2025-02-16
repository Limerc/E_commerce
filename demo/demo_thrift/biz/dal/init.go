package dal

import (
	"github.com/Limerc/E_commerce/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/Limerc/E_commerce/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
