package logic

import (
	"context"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HasHashCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHasHashCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HasHashCodeLogic {
	return &HasHashCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HasHashCodeLogic) HasHashCode(in *bucket.HashCodeReq) (*bucket.SucResp, error) {
	_, err := l.svcCtx.ObjectHashMode.FindOneByHashcode(l.ctx, in.Hashcode)

	if err != nil {
		return &bucket.SucResp{
			Success: false,
		}, nil
	}

	return &bucket.SucResp{
		Success: true,
	}, nil
}
