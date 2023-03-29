package config

import (
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/pkg/mysql"
)

// ServerConfig 服务配置信息.
type ServerConfig struct {
	Name      string       `mapstructure:"name" json:"name"`
	Port      int          `mapstructure:"port" json:"port"`
	Host      string       `mapstructure:"host" json:"host"`
	LogLevel  string       `mapstructure:"log_level" json:"log_level"`
	MysqlInfo mysql.Config `mapstructure:"mysql" json:"mysql"`
	OssConfig OssConfig    `mapstructure:"oss" json:"oss"`
}

// OssConfig oss 配置.
type OssConfig struct {
	Minio  *MinioConfig  `mapstructure:"minio" json:"minio"`
	AliYun *AliYunConfig `mapstructure:"ali_yun" json:"ali_yun"`
}

// MinioConfig minio 的配置.
type MinioConfig struct {
	Endpoint     string `mapstructure:"endpoint" json:"endpoint"`
	AccessKey    string `mapstructure:"access_key" json:"access_key"`
	AccessSecret string `mapstructure:"access_secret" json:"access_secret"`
}

// AliYunConfig 阿里云 oss 配置
type AliYunConfig struct {
	Endpoint     string `mapstructure:"endpoint" json:"endpoint"`
	AccessKey    string `mapstructure:"access_key" json:"access_key"`
	AccessSecret string `mapstructure:"access_secret" json:"access_secret"`
}
