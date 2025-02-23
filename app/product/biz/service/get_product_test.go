package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/Limerc/E_commerce/gomall/app/product/biz/dal/mysql"
	product "github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
)

func TestGetProduct_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value

	req := &product.GetProductReq{
        Id: 1,
    }
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(resp)


}
