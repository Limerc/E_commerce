package dal

import (
	"github.com/Limerc/E_commerce/gomall/app/payment/biz/dal/mysql"
	// "github.com/Limerc/E_commerce/gomall/app/payment/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
