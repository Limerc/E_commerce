package dal

import (
	"github.com/Limerc/E_commerce/gomall/app/user/biz/dal/mysql"
	"github.com/Limerc/E_commerce/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
