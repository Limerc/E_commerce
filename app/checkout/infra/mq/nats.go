package mq

import (
	"github.com/Limerc/E_commerce/gomall/app/checkout/conf"
	"github.com/nats-io/nats.go"
)

var (
	Nc *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		panic(err)
	}
}