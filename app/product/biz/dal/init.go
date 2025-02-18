package dal

import (
	"github.com/Limerc/E_commerce/gomall/app/product/biz/dal/mysql"
	"github.com/Limerc/E_commerce/gomall/app/product/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
