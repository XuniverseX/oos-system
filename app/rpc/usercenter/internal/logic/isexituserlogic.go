package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/common/xerr"

	"oos-system/app/rpc/usercenter/internal/svc"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsExitUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExitUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExitUserLogic {
	return &IsExitUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExitUserLogic) IsExitUser(in *usercenter.IsExitUserReq) (*usercenter.SucResp, error) {
	findOneByUsername, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "用户不存在 err: %v", err)
	}

	return &usercenter.SucResp{
		Success: findOneByUsername != nil,
	}, nil
}
