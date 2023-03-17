package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/auth/authservice"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/middleware"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/global"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/pkg/nacos"
)

// initOss.
func initOss() {
	// init resolver
	r, err := nacos.GetKResolve(consts.OssGroup)
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}
	// init OpenTelemetry
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(global.ServerConfig.OssSrvInfo.Name),
		provider.WithExportEndpoint(global.ServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)

	// create a new client
	c, err := authservice.NewClient(
		global.ServerConfig.OssSrvInfo.Name,
		client.WithResolver(r),                                     // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		client.WithMuxConnection(1),                                // multiplexing
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.ServerConfig.OssSrvInfo.Name}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	global.AuthClient = c
}
