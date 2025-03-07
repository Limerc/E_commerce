package main

import (
	"net"
	"time"

	"github.com/Limerc/E_commerce/gomall/app/payment/biz/dal"
	"github.com/Limerc/E_commerce/gomall/app/payment/conf"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/pkg/klog"
	//"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/joho/godotenv"
	//consul "github.com/kitex-contrib/registry-consul"
	serversuite "github.com/Limerc/E_commerce/gomall/common/serversuite"
	mtl "github.com/Limerc/E_commerce/gomall/common/mtl"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)
	
func main() {
	_ = godotenv.Load()
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	dal.Init()
	opts := kitexInit()

	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

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
