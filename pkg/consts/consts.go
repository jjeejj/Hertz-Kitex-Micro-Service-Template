package consts

const (
	JWTIssuer  = "FreeCar"
	ThirtyDays = 60 * 60 * 24 * 30
	AccountID  = "accountID"
	ID         = "id"

	ApiConfigPath  = "./server/cmd/api/config.yaml"
	AuthConfigPath = "./server/cmd/auth/config.yaml"
	OssConfigPath  = "./server/cmd/oss/config.yaml"

	ApiGroup  = "API_GROUP"
	AuthGroup = "AUTH_GROUP"
	OssGroup  = "OSS_GROUP"

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

	DefaultLicNumber = "100000000001"
	DefaultName      = "FreeCar"
	DefaultGender    = 1
	DefaultBirth     = 631152000000
)
