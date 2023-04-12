package minio

import (
	"bytes"
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	minioGo "github.com/minio/minio-go/v7"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
)

type Minio struct {
	Client *minioGo.Client
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

// PutObject 通过二进制文件 上传文件
func (minio *Minio) PutObject(ctx context.Context, req *oss.PutObjectReq) (resp *oss.PutObjectResp, err error) {
	fileReader := bytes.NewReader(req.File)
	object, err := minio.Client.PutObject(ctx, req.BucketName, req.ObjectName, fileReader, -1, minioGo.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return nil, err
	}
	klog.CtxDebugf(ctx, "object: %v", object)
	return &oss.PutObjectResp{
		Url: object.Location,
	}, nil
}
