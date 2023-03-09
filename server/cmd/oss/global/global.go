package global

import (
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
	MinioClient  *minio.Client
)
