package config

import (
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/pkg/mysql"
)

// NacosConfig 配置中心地址
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

// ServerConfig 服务配置信息
type ServerConfig struct {
	Name      string       `mapstructure:"name" json:"name"`
	Host      string       `mapstructure:"host" json:"host"`
	LogLevel  string       `mapstructure:"log_level" json:"log_level"`
	MysqlInfo mysql.Config `mapstructure:"mysql" json:"mysql"`
	OssConfig OssConfig    `mapstructure:"oss" json:"oss"`
}

// OssConfig oss 配置
type OssConfig struct {
	Minio MinioConfig `mapstructure:"minio" json:"minio"`
}

// MinioConfig minio 的配置
type MinioConfig struct {
	Endpoint     string `mapstructure:"endpoint" json:"endpoint"`
	AccessKey    string `mapstructure:"access_key" json:"access_key"`
	AccessSecret string `mapstructure:"access_secret" json:"access_secret"`
}
