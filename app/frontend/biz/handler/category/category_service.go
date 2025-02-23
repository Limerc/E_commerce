package category

import (
	"context"
	//"fmt"

	"github.com/Limerc/E_commerce/gomall/app/frontend/biz/service"
	"github.com/Limerc/E_commerce/gomall/app/frontend/biz/utils"
	category "github.com/Limerc/E_commerce/gomall/app/frontend/hertz_gen/frontend/category"

	// common "github.com/Limerc/E_commerce/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Category .
// @router /category/:category [GET]
func Category(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}	

	c.HTML(consts.StatusOK, "category", resp)
}
