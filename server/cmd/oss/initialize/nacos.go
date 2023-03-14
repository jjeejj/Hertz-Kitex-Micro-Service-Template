package initialize

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/utils"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/pkg/nacos"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
)

// InitNacos to init nacos
func InitNacos() (registry.Registry, *registry.Info) {
	r, content, err := nacos.InitNacos(consts.OssDataId, consts.OssGroup)
	if err != nil {
		klog.Fatalf("get config failed: %s", err.Error())
	}
	err = sonic.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		klog.Fatalf("nacos config failed: %s", err.Error())
	}
	// 服务端口
	if global.ServerConfig.Port == 0 {
		global.ServerConfig.Port, _ = utils.GetFreePort()
	}
	// 服务地址
	if global.ServerConfig.Host == "" {
		global.ServerConfig.Host = "0.0.0.0"
	}
	info := nacos.GetRegistryInfo(global.ServerConfig.Name, global.ServerConfig.Host, global.ServerConfig.Port)
	return r, info
}
