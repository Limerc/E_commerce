package main

import (
	"context"
	order "github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/order"
	"github.com/Limerc/E_commerce/gomall/app/order/biz/service"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// PlaceOrder implements the CartServiceImpl interface.
func (s *CartServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the CartServiceImpl interface.
func (s *CartServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}
