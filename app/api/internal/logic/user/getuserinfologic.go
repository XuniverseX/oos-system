package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
	"oos-system/app/rpc/usercenter/usercenter"
	"oos-system/common/ctxdata"
)

type GetuserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetuserInfoLogic {
	return &GetuserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserInfoLogic) GetuserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	username := ctxdata.GetUserNameFromCtx(l.ctx)

	user, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		Username: user.Username,
	}, nil
}
