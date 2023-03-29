package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
)

func InitMinio() {
	minioConfig := global.ServerConfig.OssConfig.Minio
	// 如果配置才进行初始化
	if minioConfig == nil || minioConfig.Endpoint == "" || minioConfig.AccessKey == "" || minioConfig.AccessSecret == "" {
		klog.Warn("minio config  is empty")
		return
	}
	// global mode
	var err error
	global.MinioClient, err = minio.New(minioConfig.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(minioConfig.AccessKey, minioConfig.AccessSecret, ""),
		// Secure: true,
	})
	if err != nil {
		klog.Errorf("init minio failed: %s", err.Error())
		return
	}
	klog.Info("init minio success")
}
