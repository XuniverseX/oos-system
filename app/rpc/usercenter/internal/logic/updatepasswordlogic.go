package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/common/tool"
	"oos-system/common/xerr"
	"time"

	"oos-system/app/rpc/usercenter/internal/svc"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(in *usercenter.LoginReq) (*usercenter.SucResp, error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdatePassword findOneByUsername db err , username : %s , err : %+v", in.Username, err)
	}

	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("UpdatePassword user no exists"), " username : %s", in.Username)
	}

	if len(in.Password) > 0 {
		user.Password = tool.Md5ByString(in.Password)
	}

	user.UpdateTime = time.Now()

	updateErr := l.svcCtx.UserModel.Update(l.ctx, user)
	if updateErr != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " UpdatePassword Update db  username : %s , err : %+v", in.Username, updateErr)
	}

	return &usercenter.SucResp{
		Success: updateErr == nil,
	}, nil
}
