package mq

import (
	"github.com/Limerc/E_commerce/gomall/app/email/conf"
	"github.com/nats-io/nats.go"
)

var (
	Nc *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		print("Error connecting to NATS: " + err.Error())
		panic(err)
	}
}