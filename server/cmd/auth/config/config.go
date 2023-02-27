package config

import (
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/pkg/mysql"
)

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type WXConfig struct {
	AppId     string `mapstructure:"app_id" json:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret"`
}

type ServerConfig struct {
	Name      string       `mapstructure:"name" json:"name"`
	Host      string       `mapstructure:"host" json:"host"`
	LogLevel  string       `mapstructure:"log_level" json:"log_level"`
	MysqlInfo mysql.Config `mapstructure:"mysql" json:"mysql"`
	OtelInfo  OtelConfig   `mapstructure:"otel" json:"otel"`
	WXInfo    WXConfig     `mapstructure:"wx_config" json:"wx_config"`
}
