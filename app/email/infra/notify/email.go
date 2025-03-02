package notify

import (
	"github.com/kr/pretty"
	//"github.com/Limerc/E_commerce/gomall/rpc_gen/rpc/email"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/email"
)

type NoopEmail struct{}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	pretty.Printf("%v\n", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}