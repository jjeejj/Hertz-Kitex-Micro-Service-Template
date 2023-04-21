package minio

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	minioGo "github.com/minio/minio-go/v7"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
)

type Minio struct {
	Client *minioGo.Client
}

// PreSignedPutObjectUrl 获取 put 上传文件的签名地址
func (minio *Minio) PreSignedPutObjectUrl(ctx context.Context, req *oss.PreSignedPutObjectUrlReq) (*oss.PreSignedPutObjectUrlResp, error) {
	objUrl, err := minio.Client.PresignedPutObject(ctx, req.BucketName, req.ObjectName, time.Second*time.Duration(req.Expiry))
	if err != nil {
		klog.Errorf("MinioClient PreSignedPutObject error: %v", err)
		return nil, err
	}
	klog.Debugf("MinioClient PreSignedPutObject success %v", objUrl)
	preSignedPutObjectUrl := &oss.PreSignedPutObjectUrlResp{
		PreSignedUrl: objUrl.String(),
		ResourceUrl:  fmt.Sprintf("%s://%s/%s/%s", global.ServerConfig.OssConfig.Minio.Scheme, global.ServerConfig.OssConfig.Minio.Endpoint, req.BucketName, url.PathEscape(req.ObjectName)),
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
