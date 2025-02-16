package service

import (
	"context"
	"testing"
	pbapi "github.com/Limerc/E_commerce/gomall/demo/demo_proto/kitex_gen/pbapi"
)

func TestEcho_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoService(ctx)
	// init req and assert value

	req := &pbapi.Request{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
