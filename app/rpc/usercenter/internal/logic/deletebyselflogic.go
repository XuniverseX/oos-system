package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"oos-system/app/rpc/usercenter/internal/svc"
	"oos-system/app/rpc/usercenter/usercenterclient"
	"oos-system/common/xerr"
)

type DeleteBySelfLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBySelfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBySelfLogic {
	return &DeleteBySelfLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBySelfLogic) DeleteBySelf(in *usercenterclient.DelReq) (*usercenterclient.SucResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DeleteBySelf findOneByUsername db err , username : %s , err : %+v", in.Username, err)
	}

	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("DeleteBySelf user no exists"), " username : %s", in.Username)
	}
	delErr := l.svcCtx.UserModel.Delete(l.ctx, user.Id)

	if delErr != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " DeleteBySelf db err: username : %s , err : %+v", in.Username, delErr)
	}

	// 不搞逻辑删除了 冗余得很
	//user.Deleted = 1
	//user.UpdateTime = time.Now()
	//
	//updateErr := l.svcCtx.UserModel.Update(l.ctx, user)
	//if updateErr != nil {
	//	return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " DeleteBySelf Update db  username : %s , err : %+v", in.Username, updateErr)
	//}

	return &usercenterclient.SucResp{
		Success: delErr == nil,
	}, nil
}
