package email

import (
	"context"
	email "github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/email"

	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/email/emailservie"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() emailservie.Client
	Service() string
	Send(ctx context.Context, Req *email.EmailReq, callOptions ...callopt.Option) (r *email.EmailResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := emailservie.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient emailservie.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() emailservie.Client {
	return c.kitexClient
}

func (c *clientImpl) Send(ctx context.Context, Req *email.EmailReq, callOptions ...callopt.Option) (r *email.EmailResp, err error) {
	return c.kitexClient.Send(ctx, Req, callOptions...)
}
