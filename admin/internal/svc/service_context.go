package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"iot-platform/admin/internal/config"
	"iot-platform/model"
	"iot-platform/user/rpc/types/user"
	"iot-platform/user/rpc/userclient"
)

type ServiceContext struct {
	DB       *gorm.DB
	Config   config.Config
	RpcUser  userclient.User
	AuthUser *user.UserAuthReply
}

func NewServiceContext(c config.Config) *ServiceContext {
	model.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      model.DB,
		RpcUser: userclient.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
