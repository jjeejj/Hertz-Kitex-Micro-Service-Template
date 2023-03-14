package main

import (
	"net"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/log"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/middleware"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"

	oss "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss/ossservice"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/initialize"
)

func main() {
	// initialization
	r, info := initialize.InitNacos()
	log.InitKLogger(consts.KlogFilePath, global.ServerConfig.LogLevel)
	initialize.InitDB()

	// init minio
	initialize.InitMinio()

	// Create new server.
	srv := oss.NewServer(new(OssServiceImpl),
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(global.ServerConfig.Host, strconv.Itoa(global.ServerConfig.Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.ServerConfig.Name}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithMuxTransport(),
	)

	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
