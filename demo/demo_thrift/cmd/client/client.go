package main

import (
	"context"
	"fmt"

	"github.com/Limerc/E_commerce/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/Limerc/E_commerce/gomall/demo/demo_thrift/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)


func main() {
	cli, err := echo.NewClient("demo_thrift", client.WithHostPorts("localhost:8888"),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),  // 配置thrift对应的元信息处理器
		client.WithTransportProtocol(transport.TTHeader),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_thrift_client",
		}),
	) // 服务名称、服务地址
	if err != nil {
		panic(err)
	}

	res, err := cli.Echo(context.Background(), &api.Request{
		Message: "hello kitex",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v",res)
}