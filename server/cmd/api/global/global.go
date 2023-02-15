package global

import (
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/config"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/kitex_gen/auth/authservice"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/kitex_gen/car/carservice"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/kitex_gen/profile/profileservice"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/kitex_gen/trip/tripservice"
)

var (
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}

	AuthClient    authservice.Client
	CarClient     carservice.Client
	ProfileClient profileservice.Client
	TripClient    tripservice.Client
)
