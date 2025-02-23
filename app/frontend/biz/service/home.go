package service

import (
	"context"

	common "github.com/Limerc/E_commerce/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/Limerc/E_commerce/gomall/app/frontend/infra/rpc"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	// // 入参为产品列表
	// var resp = make(map[string]any)
	// items := []map[string]any{
	// 	{"Name":"藕粉哪吒","Price":59,"Picture":"/static/image/product/Product1.png"},
	// 	{"Name":"藕粉敖丙","Price":59,"Picture":"/static/image/product/Product2.png"},
	// 	{"Name":"敖丙版哪吒","Price":59,"Picture":"/static/image/product/Product3.png"},
	// 	{"Name":"乖巧敖丙","Price":69,"Picture":"/static/image/product/Product4.png"},
	// 	{"Name":"捣蛋哪吒","Price":69,"Picture":"/static/image/product/Product5.png"},
	// 	{"Name":"牵手敖丙","Price":79,"Picture":"/static/image/product/Product6.png"},
	// 	{"Name":"牵手哪吒","Price":79,"Picture":"/static/image/product/Product7.png"},
	// }
	// resp["Title"] = "Surrrounding NeZha"
	// resp["items"] = items
	// return resp, nil
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
