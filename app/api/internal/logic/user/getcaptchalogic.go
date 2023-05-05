package user

import (
	"context"
	"github.com/jinzhu/copier"
	"oos-system/app/rpc/usercenter/usercenter"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetcaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetcaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetcaptchaLogic {
	return &GetcaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetcaptchaLogic) Getcaptcha(req *types.GetCaptchaReq) (resp *types.GetCaptchaResp, err error) {
	// todo: add your logic here and delete this line
	if len(req.Id) != 0 {
		_, err := l.svcCtx.UsercenterRpc.DelOldCaptcha(l.ctx, &usercenter.DelOldCaptchaReq{
			Id: req.Id,
		})

		if err != nil {
			return nil, err
		}
	}
	captcha, err := l.svcCtx.UsercenterRpc.GetCaptcha(l.ctx, &usercenter.GetCaptchaReq{})
	if err != nil {
		return nil, err
	}

	var res types.GetCaptchaResp
	_ = copier.Copy(&res, captcha)

	return &res, nil
}
