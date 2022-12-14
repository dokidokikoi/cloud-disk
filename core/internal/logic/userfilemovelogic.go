package logic

import (
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	parentDate := new(models.UserRepository)
	l.svcCtx.Engine.ShowSQL(true)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).Get(parentDate)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件夹不存在")
	}

	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(models.UserRepository{
		ParentId: int64(parentDate.Id),
	})

	return
}
