package global

import (
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/config"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/kitex_gen/blob/blobservice"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig

	BlobClient blobservice.Client
)
