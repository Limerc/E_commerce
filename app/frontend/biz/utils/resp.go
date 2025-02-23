package utils

import (
	"context"

	// "github.com/Limerc/E_commerce/gomall/app/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	// "github.com/hertz-contrib/sessions"
	"github.com/Limerc/E_commerce/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/Limerc/E_commerce/gomall/app/frontend/utils"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/cart"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any{
	userId := frontendUtils.GetUserIdFromCtx(ctx)
	content["user_id"] = userId

	if userId > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: uint32(frontendUtils.GetUserIdFromCtx(ctx)),
		})
		if err == nil && cartResp.Items != nil {
			content["cart_num"] = len(cartResp.Items)
		}
	}
	return content
}