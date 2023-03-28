package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
)

func InitMinio() {
	minioConfig := global.ServerConfig.OssConfig.Minio
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
