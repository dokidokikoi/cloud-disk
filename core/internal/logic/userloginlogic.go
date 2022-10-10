package logic

import (
	"context"
	"errors"

	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	user := models.FindUserByNameAndPwd(req.Name, req.Password)
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}

	token, err := helper.GenerateToke(user.Id, user.Identity, user.Name, int64(define.TokenExpire))
	if err != nil {
		return nil, err
	}

	// 生成用于刷新token的token
	refreshToken, err := helper.GenerateToke(user.Id, user.Identity, user.Name, int64(define.RefreshTokenExpire))
	if err != nil {
		return nil, err
	}

	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
