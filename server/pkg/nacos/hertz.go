package nacos

import (
	"log"
	"net"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
)

func GetHRegistryInfo(serverName, serverHost string, serverPort int) (registry.Registry, *registry.Info) {
	if serverName == "" || serverHost == "" || serverPort == 0 {
		log.Fatal("serverName , serverHost , serverPort is empty")
	}
	registryClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId:         configInfo.Namespace,
				TimeoutMs:           5000,
				NotLoadCacheAtStart: true,
				LogDir:              consts.NacosLogDir,
				CacheDir:            consts.NacosCacheDir,
				LogLevel:            consts.NacosLogLevel,
			},
			ServerConfigs: []constant.ServerConfig{
				{
					IpAddr: configInfo.Host,
					Port:   configInfo.Port,
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("get registryClient failed: %s", err.Error())
	}
	if err != nil {
		log.Fatalf("get registryClient failed: %s", err.Error())
	}
	r := nacos.NewNacosRegistry(registryClient, nacos.WithRegistryGroup(configInfo.Group))
	sf, err := snowflake.NewNode(2)
	if err != nil {
		log.Fatalf("generate snowflake failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: serverName,
		Addr:        utils.NewNetAddr(consts.TCP, net.JoinHostPort(serverHost, strconv.Itoa(serverPort))),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
		Weight: registry.DefaultWeight,
	}

	return r, info
}
