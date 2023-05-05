package ctxdata

import (
	"context"
	"fmt"
)

// CtxKeyJwtUserName get username from ctx
var CtxKeyJwtUserName = "username"

func GetUserNameFromCtx(ctx context.Context) string {
	var username string
	username = fmt.Sprintf("%v", ctx.Value(CtxKeyJwtUserName))

	return username
}
