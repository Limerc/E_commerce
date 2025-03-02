package rpc

import (
	"sync"

	"github.com/Limerc/E_commerce/gomall/app/cart/conf"
	cartUtils "github.com/Limerc/E_commerce/gomall/app/cart/utils"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	// consul "github.com/kitex-contrib/registry-consul"
	clientsuite "github.com/Limerc/E_commerce/gomall/common/clientsuite"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once // 保证初始化只执行一次
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
