package main

import (
	"context"
	email "github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/email"
	"github.com/Limerc/E_commerce/gomall/app/email/biz/service"
)

// EmailServieImpl implements the last service interface defined in the IDL.
type EmailServieImpl struct{}

// Send implements the EmailServieImpl interface.
func (s *EmailServieImpl) Send(ctx context.Context, req *email.EmailReq) (resp *email.EmailResp, err error) {
	resp, err = service.NewSendService(ctx).Run(req)

	return resp, err
}
