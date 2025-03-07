package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/Limerc/E_commerce/gomall/app/user/biz/dal/mysql"
	user "github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "test3@test.com",
		Password:        "GADFGADVdf",
		PasswordConfirm: "GADFGADVdf",
	}

	resp, err := s.Run(req)
	fmt.Println(resp, err)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
