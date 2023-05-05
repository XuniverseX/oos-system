package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelHashCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelHashCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelHashCodeLogic {
	return &DelHashCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelHashCodeLogic) DelHashCode(in *bucket.HashCodeReq) (*bucket.SucResp, error) {

	findOneByHashcode, err := l.svcCtx.ObjectHashMode.FindOneByHashcode(l.ctx, in.Hashcode)

	if err == sqlx.ErrNotFound {
		return &bucket.SucResp{
			Success: true,
		}, nil
	}

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find hashcode db err: %v", err)
	}

	err = l.svcCtx.ObjectHashMode.Delete(l.ctx, findOneByHashcode.Id)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete hashcode db err: %v", err)
	}

	return &bucket.SucResp{
		Success: err == nil,
	}, nil
}
