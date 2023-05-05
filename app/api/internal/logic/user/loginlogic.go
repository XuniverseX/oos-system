package user

import (
	"context"
	"github.com/jinzhu/copier"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	loginResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	var res types.LoginResp
	_ = copier.Copy(&res, loginResp)

	return &res, nil
}
