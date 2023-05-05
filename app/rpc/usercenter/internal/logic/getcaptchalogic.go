package logic

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/pkg/errors"
	"oos-system/app/rpc/usercenter/internal/utils"
	"oos-system/app/rpc/usercenter/usercenterclient"
	"oos-system/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"oos-system/app/rpc/usercenter/internal/svc"
)

var ErrGetCaptcha = xerr.NewErrMsg("获取验证码失败")

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaLogic) GetCaptcha(in *usercenterclient.GetCaptchaReq) (*usercenterclient.GetCaptchaResp, error) {
	// todo: add your logic here and delete this line
	//配置RedisStore RedisStore实现base64Captcha.Store的接口
	var store base64Captcha.Store = utils.CaptchaRedisStore{}

	digitType := &base64Captcha.DriverDigit{
		Height:   50,
		Width:    100,
		Length:   4,
		MaxSkew:  0.45,
		DotCount: 80,
	}
	var driver base64Captcha.Driver = digitType

	c := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := c.Generate()

	if err != nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "验证码获取失败")
	}

	return &usercenterclient.GetCaptchaResp{
		Id:          id,
		ImgbaseCode: b64s,
	}, nil
}
