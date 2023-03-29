package initialize

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/translate/global"
)

func InitAliYun() {
	aliYunConfig := global.ServerConfig.AliYun
	if aliYunConfig == nil || aliYunConfig.AccessId == "" || aliYunConfig.AccessSecret == "" {
		klog.Warn("AliYun config is empty")
		return
	}
	// global mode
	var err error
	global.AliYunClient, err = alimt.NewClientWithAccessKey(aliYunConfig.Region, aliYunConfig.AccessId, aliYunConfig.AccessSecret)
	// 判断是否设置代理
	if global.ServerConfig.AliYun.Proxy != "" {
		global.AliYunClient.SetHttpsProxy(global.ServerConfig.AliYun.Proxy)
		global.AliYunClient.SetHttpProxy(global.ServerConfig.AliYun.Proxy)
	}
	if err != nil {
		klog.Errorf("init ali yun translate fail: %s", err.Error())
		return
	}
	klog.Info("init ali yun translate  success")
}
