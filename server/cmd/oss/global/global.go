package global

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/minio/minio-go/v7"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/config"
)

var (
	ServerConfig config.ServerConfig
	MinioClient  *minio.Client
	AliYunClient *aliOss.Client
)
