package service

import (
	"context"

	// common "github.com/Limerc/E_commerce/gomall/app/frontend/hertz_gen/frontend/common"
	product "github.com/Limerc/E_commerce/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/Limerc/E_commerce/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Query,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items" : products.Results,
		"q": 	 req.Query,
	}, nil
}
