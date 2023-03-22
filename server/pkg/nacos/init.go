package nacos

import (
	"log"
	"os"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
)

var configInfo *ConfigInfo

// InitNacos 根据不同的环境初始化 nacos.
// 返回对应  DataId 和 Group 下的配置数据.
func InitNacos(dataId, group string) (string, error) {
	v := viper.New()
	// E
	v.SetConfigFile(consts.NacosConfigPath)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("read viper config failed: %s", err.Error())
	}
	config := Config{}
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("unmarshal err failed: %s", err.Error())
	}

	// 读取环境变量 HOST_ENV
	// 存在 dev | test | prod
	envValue := os.Getenv(consts.EnvKey)
	switch envValue {
	case consts.DevEnv:
		configInfo = config.Dev
	case consts.TestEnv:
		configInfo = config.Test
	case consts.ProdEnv:
		configInfo = config.Prod
	default:
		envValue = consts.DevEnv
		configInfo = config.Dev // 默认是 dev 环境
	}
	log.Printf("HOST_ENV: %v, active Config Info: %#v", envValue, configInfo)
	// Read configuration information from nacos
	sc := []constant.ServerConfig{
		{
			IpAddr: configInfo.Host,
			Port:   configInfo.Port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         configInfo.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              consts.NacosLogDir,
		CacheDir:            consts.NacosCacheDir,
		LogLevel:            consts.NacosLogLevel,
		Username:            configInfo.User,
		Password:            configInfo.Password,
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		log.Fatalf("create config client failed: %s", err.Error())
	}
	configInfo.DataId = dataId
	configInfo.Group = group
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		log.Fatalf("get config content failed: %s", err.Error())
	}

	return content, err
}
