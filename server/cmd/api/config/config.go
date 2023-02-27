package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ServerConfig struct {
	Name        string       `mapstructure:"name" json:"name"`
	Host        string       `mapstructure:"host" json:"host"`
	Port        int          `mapstructure:"port" json:"port"`
	LogLevel    string       `mapstructure:"log_level" json:"log_level"`
	JWTInfo     JWTConfig    `mapstructure:"jwt" json:"jwt"`
	OtelInfo    OtelConfig   `mapstructure:"otel" json:"otel"`
	AuthSrvInfo RPCSrvConfig `mapstructure:"auth_srv" json:"auth_srv"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
