package config

// ServerConfig 服务配置信息.
type ServerConfig struct {
	Name           string         `mapstructure:"name" json:"name"`
	Port           int            `mapstructure:"port" json:"port"`
	Host           string         `mapstructure:"host" json:"host"`
	LogLevel       string         `mapstructure:"log_level" json:"log_level"`
	MqServerConfig MqServerConfig `mapstructure:"mq_server_config" json:"mq_server_config"`
}

type MqServerConfig struct {
	Host string `mapstructure:"host"`
}
