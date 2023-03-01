package global

import (
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/auth/authservice"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}

	AuthClient authservice.Client
)
