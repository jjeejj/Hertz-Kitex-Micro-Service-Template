package minio

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/initialize"
)

func init() {
	initialize.InitNacos(0)
	initialize.InitMinio()
}

func TestPreSignedPutObjectUrl(t *testing.T) {
	GetOssImpl := &Minio{
		Client: global.MinioClient,
	}
	resp, err := GetOssImpl.PreSignedPutObjectUrl(context.Background(), &oss.PreSignedPutObjectUrlReq{
		BucketName: "test",
		ObjectName: "test.png",
		Expiry:     1000,
	})
	if err != nil {
		klog.Errorf("err: %v", err)
		return
	}
	klog.Infof("resp 	%v", resp)
}
