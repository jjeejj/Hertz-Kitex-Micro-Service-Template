package consts

const (
	JWTIssuer  = "FreeCar"
	ThirtyDays = 60 * 60 * 24 * 30
	AccountID  = "accountID"
	ID         = "id"

	OssConfigPath   = "./server/cmd/oss/config.yaml"
	NacosConfigPath = "./config/nacos_config.yaml"

	AuthGroup = "AUTH_GROUP"

	NacosLogDir   = "tmp/nacos/log"
	NacosCacheDir = "tmp/nacos/cache"
	NacosLogLevel = "debug"

	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	MySqlDSN    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MongoURI    = "mongodb://%s:%s@%s:%d"
	RabbitMqURI = "amqp://%s:%s@%s:%d/"

	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"

	TCP = "tcp"

	FreePortAddress = "localhost:0"
)

// 环境相关的.
const (
	EnvKey  = "HOST_ENV" // 配置的环境变量
	DevEnv  = "dev"      // dev 环境
	TestEnv = "test"
	ProdEnv = "prod"
)

// 和 nacos 相关的微服务 常量.
const (
	OssGroup  = "OSS_GROUP"
	OssDataId = "oss_srv"

	ApiGroup  = "API_GROUP"
	ApiDataId = "api"
)
