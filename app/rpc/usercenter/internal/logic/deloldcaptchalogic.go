package logic

import (
	"context"
	"oos-system/app/rpc/usercenter/internal/svc"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOldCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOldCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOldCaptchaLogic {
	return &DelOldCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOldCaptchaLogic) DelOldCaptcha(in *usercenter.DelOldCaptchaReq) (*usercenter.SucResp, error) {
	// todo: add your logic here and delete this line
	del, err := l.svcCtx.RedisClient.Del("captcha:" + in.Id)
	if err != nil {
		return nil, err
	}

	return &usercenter.SucResp{
		Success: del == 1,
	}, nil
}
