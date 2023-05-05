package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line

	register, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Username:   req.Username,
		Password:   req.Password,
		Captcha:    req.Captcha,
		UniqueCode: req.UniqueCode,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var res types.RegisterResp
	_ = copier.Copy(&res, register)

	return &res, nil
}
