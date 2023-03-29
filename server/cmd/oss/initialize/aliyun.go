package initialize

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
)

func InitAliYun() {
	aliYunConfig := global.ServerConfig.OssConfig.AliYun
	// 如果配置才进行初始化
	if aliYunConfig == nil || aliYunConfig.Endpoint == "" || aliYunConfig.AccessKey == "" || aliYunConfig.AccessSecret == "" {
		klog.Warn("ali yun config  is empty")
		return
	}
	// global mode
	var err error
	global.AliYunClient, err = aliOss.New(aliYunConfig.Endpoint, aliYunConfig.AccessKey, aliYunConfig.AccessSecret)
	if err != nil {
		klog.Errorf("init ali yun oss  failed: %s", err.Error())
		return
	}
	klog.Info("init ali yun oss  success")
}
