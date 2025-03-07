package main

import (
	"net"
	"time"

	"github.com/Limerc/E_commerce/gomall/app/checkout/conf"
	"github.com/Limerc/E_commerce/gomall/app/checkout/infra/mq"
	"github.com/Limerc/E_commerce/gomall/app/checkout/infra/rpc"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/pkg/klog"
	//"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	//consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	serversuite "github.com/Limerc/E_commerce/gomall/common/serversuite"
	mtl "github.com/Limerc/E_commerce/gomall/common/mtl"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)
	
func main() {
	opts := kitexInit()
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	rpc.InitClient()
	mq.Init()
	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))

	// r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	// if err != nil {
	// 	klog.Fatal(err)
	// }
	// opts = append(opts, server.WithServiceAddr(addr),server.WithRegistry(r))

	// // service info
	// opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
	// 	ServiceName: conf.GetConf().Kitex.Service,
	// }))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
