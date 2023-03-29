package aliyun

import (
	"context"

	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
)

type AliYun struct {
	Client *aliOss.Client
}

// PreSignedPutObjectUrl 阿里云生成预签名的链接
func (ali *AliYun) PreSignedPutObjectUrl(ctx context.Context, req *oss.PreSignedPutObjectUrlReq) (*oss.PreSignedPutObjectUrlResp, error) {
	bucket, err := ali.Client.Bucket(req.BucketName)
	if err != nil {
		klog.CtxErrorf(ctx, "AliYun Oss Bucket error: %v", err)
		return nil, err
	}
	url, err := bucket.SignURL(req.ObjectName, aliOss.HTTPPut, int64(req.Expiry))
	if err != nil {
		klog.CtxErrorf(ctx, "AliYun os  SignURL error: %v", err)
		return nil, err
	}
	return &oss.PreSignedPutObjectUrlResp{
		PreSignedUrl: url,
	}, nil
}
