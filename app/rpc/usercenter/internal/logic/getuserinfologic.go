package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/model/usermodel"
	"oos-system/common/xerr"

	"oos-system/app/rpc/usercenter/internal/svc"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *usercenter.UserInfoReq) (*usercenter.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != usermodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , username:%s , err:%v", in.Username, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "username:%s", in.Username)
	}

	return &usercenter.UserInfoResp{
		Username: user.Username,
	}, nil
}
