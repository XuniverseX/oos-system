package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/model/objecthashmodel"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHashCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHashCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHashCodeLogic {
	return &AddHashCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddHashCodeLogic) AddHashCode(in *bucket.HashCodeReq) (*bucket.SucResp, error) {

	objectHash := new(objecthashmodel.ObjectHash)

	objectHash.Hashcode = in.Hashcode

	_, err := l.svcCtx.ObjectHashMode.Insert(l.ctx, objectHash)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert hashCode db err: %v", err)
	}

	return &bucket.SucResp{
		Success: err == nil,
	}, nil
}
