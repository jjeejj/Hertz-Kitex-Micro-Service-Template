package initialize

import (
	"net"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	nacos "github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"

	pkgutils "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/utils"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
)

// InitNacos to init nacos
func InitNacos() (registry.Registry, *registry.Info) {
	v := viper.New()
	v.SetConfigFile(consts.OssConfigPath)
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	klog.Infof("Config Info: %v", global.NacosConfig)

	// Read configuration information from nacos
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              consts.NacosLogDir,
		CacheDir:            consts.NacosCacheDir,
		LogLevel:            consts.NacosLogLevel,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		klog.Fatalf("create config client failed: %s", err.Error())
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})
	if err != nil {
		klog.Fatalf("get config failed: %s", err.Error())
	}

	err = sonic.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		klog.Fatalf("nacos config failed: %s", err.Error())
	}

	registryClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	r := nacos.NewNacosRegistry(registryClient, nacos.WithGroup(global.NacosConfig.Group))

	sf, err := snowflake.NewNode(2)
	if err != nil {
		klog.Fatalf("generate service name failed: %s", err.Error())
	}
	if global.ServerConfig.Port == 0 {
		global.ServerConfig.Port, _ = pkgutils.GetFreePort()
	}
	info := &registry.Info{
		ServiceName: global.ServerConfig.Name,
		Addr:        utils.NewNetAddr(consts.TCP, net.JoinHostPort(global.ServerConfig.Host, strconv.Itoa(global.ServerConfig.Port))),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
