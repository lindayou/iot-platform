package logic

import (
	"context"
	"errors"
	"iot-platform/helper"
	"iot-platform/model"

	"iot-platform/user/internal/svc"
	"iot-platform/user/internal/types"

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

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserLoginReply)
	uc := new(model.UserBasic)
	err = l.svcCtx.DB.Where("name = ? and password =?", req.Name, helper.Md5(req.Password)).First(uc).Error
	if err != nil {
		logx.Error("[DB ERROR :]", err)
		err = errors.New("用户名或密码错误")
		return nil, err
	}
	token := helper.GenerateToken(uc.ID, uc.Name, uc.Identity, 3600*24*30)
	resp.Token = token

	return
}
