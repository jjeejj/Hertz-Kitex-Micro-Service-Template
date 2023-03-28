package minio

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
)

type Minio struct {
	Client *minio.Client
}

// PreSignedPutObjectUrl 获取 put 上传文件的签名地址
func (minio *Minio) PreSignedPutObjectUrl(ctx context.Context, req *oss.PreSignedPutObjectUrlReq) (*oss.PreSignedPutObjectUrlResp, error) {
	url, err := minio.Client.PresignedPutObject(ctx, req.BucketName, req.ObjectName, time.Second*time.Duration(req.Expiry))
	if err != nil {
		klog.Errorf("MinioClient PreSignedPutObject error: %v", err)
		return nil, err
	}
	klog.Debugf("MinioClient PreSignedPutObject success %v", url)
	preSignedPutObjectUrl := &oss.PreSignedPutObjectUrlResp{
		PreSignedUrl: url.String(),
	}
	return preSignedPutObjectUrl, nil
}
