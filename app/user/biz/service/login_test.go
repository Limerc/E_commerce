package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/Limerc/E_commerce/gomall/app/user/biz/dal/mysql"
	user "github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestLogin_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:           "test3@test.com",
		Password:        "GADFGADVd",
	}

	resp, err := s.Run(req)
	fmt.Println(resp, err)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
