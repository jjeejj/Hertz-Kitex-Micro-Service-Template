package main

import (
	"context"
	"net"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	auth "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/auth/authservice"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/log"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/middleware"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/global"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/initialize"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/tool"
)

func main() {
	// initialization
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitNacos(Port)
	log.InitKLogger(consts.KlogFilePath, global.ServerConfig.LogLevel)
	initialize.InitDB()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(global.ServerConfig.Name),
		provider.WithExportEndpoint(global.ServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	impl := new(AuthServiceImpl)
	impl.OpenIDResolver = &tool.AuthServiceImpl{
		AppID:     global.ServerConfig.WXInfo.AppId,
		AppSecret: global.ServerConfig.WXInfo.AppSecret,
	}
	// Create new server.
	srv := auth.NewServer(impl,
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.ServerConfig.Name}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)

	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
