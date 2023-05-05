package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/model/usermodel"
	"oos-system/common/tool"
	"oos-system/common/xerr"

	"oos-system/app/rpc/usercenter/internal/svc"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *usercenter.LoginReq) (*usercenter.LoginResp, error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != usermodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询用户信息失败，username:%s,err:%v", in.Username, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "username:%s", in.Username)
	}

	if user.Deleted == 1 {
		return nil, errors.Wrapf(xerr.NewErrMsg("账号已注销"), "username: %s", in.Username)
	}

	if !(tool.Md5ByString(in.Password) == user.Password) {
		return nil, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		Username: user.Username,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken username : %s", user.Username)
	}

	return &usercenter.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
