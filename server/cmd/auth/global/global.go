package global

import (
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
