package svc

import (
	"gorm.io/gorm"
	"iot-platform/model"
	"iot-platform/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	model.NewDB()
	return &ServiceContext{
		Config: c,
		DB:     model.DB,
	}
}
