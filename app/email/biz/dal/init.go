package dal

import (
	"github.com/Limerc/E_commerce/gomall/app/email/biz/dal/mysql"
	"github.com/Limerc/E_commerce/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
