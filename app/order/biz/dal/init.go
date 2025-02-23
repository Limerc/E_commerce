package dal

import (
	"github.com/Limerc/E_commerce/gomall/app/order/biz/dal/mysql"
	"github.com/Limerc/E_commerce/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
