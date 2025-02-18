package rpc

import (
	"sync"

	"github.com/Limerc/E_commerce/gomall/app/frontend/conf"
	frontendUtils "github.com/Limerc/E_commerce/gomall/app/frontend/utils"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/user/userservice"
	// "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)


var (
	UserClient userservice.Client

	once sync.Once   // 保证初始化只执行一次
)

func Init(){
	once.Do(func() {
		iniUserClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}