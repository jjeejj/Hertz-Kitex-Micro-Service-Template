package global

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/translate/config"
)

var (
	ServerConfig config.ServerConfig
	AliYunClient *alimt.Client
)
