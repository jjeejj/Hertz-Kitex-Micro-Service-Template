package config

// ServerConfig 服务配置信息.
type ServerConfig struct {
	Name     string        `mapstructure:"name" json:"name"`
	Port     int           `mapstructure:"port" json:"port"`
	Host     string        `mapstructure:"host" json:"host"`
	LogLevel string        `mapstructure:"log_level" json:"log_level"`
	AliYun   *AliYunConfig `mapstructure:"ali_yun" json:"ali_yun"`
}

// AliYunConfig 阿里云翻译配置
type AliYunConfig struct {
	AccessId     string `mapstructure:"access_id" json:"access_id"`
	AccessSecret string `mapstructure:"access_secret" json:"access_secret"`
	Region       string `mapstructure:"region" json:"region"`
}
