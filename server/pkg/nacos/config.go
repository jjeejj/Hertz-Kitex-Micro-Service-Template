package nacos

// Config 各个环境的信息.
type Config struct {
	Local *ConfigInfo `json:"local" mapstructure:"local"`
	Dev   *ConfigInfo `json:"dev" mapstructure:"dev"`
	Test  *ConfigInfo `json:"test" mapstructure:"test"`
	Prod  *ConfigInfo `json:"prod" mapstructure:"prod"`
}

// ConfigInfo 具体的配置信息.
type ConfigInfo struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
	Weight    int    `mapstructure:"weight"`
}
