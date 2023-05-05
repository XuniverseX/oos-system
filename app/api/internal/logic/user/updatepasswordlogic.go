package user

import (
	"context"
	"github.com/jinzhu/copier"
	"oos-system/app/rpc/usercenter/usercenter"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.LoginReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	updatePassword, err := l.svcCtx.UsercenterRpc.UpdatePassword(l.ctx, &usercenter.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	var res types.RegisterResp
	_ = copier.Copy(&res, updatePassword)

	return &res, nil
}
