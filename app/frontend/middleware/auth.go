package middleware

import (
	"context"

	frontendUtils "github.com/Limerc/E_commerce/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIdKey string
const SessionUserId SessionUserIdKey = "user_id"

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext){
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext){
		s := sessions.Default(c)
		userId := s.Get("user_id") // 从Cookie中获取user_id
		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))  // 跳转到登录页面
			c.Abort()    // 打断后续操作
			return
		}
		c.Next(ctx)      // 获取到，走下一个逻辑
	}
}

