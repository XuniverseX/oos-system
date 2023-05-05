package logic

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/pkg/errors"
	"oos-system/app/rpc/usercenter/internal/utils"
	"oos-system/common/tool"
	"oos-system/common/xerr"

	"oos-system/app/rpc/model/usermodel"
	"oos-system/app/rpc/usercenter/internal/svc"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrCaptchaError = xerr.NewErrMsg("验证码错误")
var ErrUserAlreadyRegisterError = xerr.NewErrMsg("用户名已经被注册")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.SucResp, error) {
	// todo: add your logic here and delete this line
	username, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != usermodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username:%s,err:%v", in.Username, err)
	}
	if username != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register username exists username:%s,err:%v", in.Username, err)
	}

	var store base64Captcha.Store = utils.CaptchaRedisStore{}
	id := "captcha:" + in.UniqueCode

	if store.Verify(id, in.Captcha, true) {
		return nil, errors.Wrapf(ErrCaptchaError, "Captcha not verify：err:%v", err)
	}

	user := new(usermodel.User)

	user.Username = in.Username

	if len(in.Password) > 0 {
		user.Password = tool.Md5ByString(in.Password)
	}

	_, insertErr := l.svcCtx.UserModel.Insert(l.ctx, user)
	if insertErr != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", insertErr, user)
	}

	return &usercenter.SucResp{
		Success: insertErr == nil,
	}, nil
}
