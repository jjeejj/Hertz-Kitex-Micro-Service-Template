package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/biz/errno"
)

func Recovery() app.HandlerFunc {
	return recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, utils.H{
				"code":    errno.BadRequest,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		},
	))
}
