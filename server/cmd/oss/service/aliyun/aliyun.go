package aliyun

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
)

type AliYun struct {
	Client *aliOss.Client
}

// PreSignedPutObjectUrl 阿里云生成预签名的链接.
// 这里生成签名的时候没有 指定 Content-Type 所以前端上传的时候也不能指定 Content-Type.
func (ali *AliYun) PreSignedPutObjectUrl(ctx context.Context, req *oss.PreSignedPutObjectUrlReq) (*oss.PreSignedPutObjectUrlResp, error) {
	bucket, err := ali.Client.Bucket(req.BucketName)
	if err != nil {
		klog.CtxErrorf(ctx, "AliYun Oss Bucket error: %v", err)
		return nil, err
	}
	objUrl, err := bucket.SignURL(req.ObjectName, aliOss.HTTPPut, int64(req.Expiry))
	if err != nil {
		klog.CtxErrorf(ctx, "AliYun os  SignURL error: %v", err)
		return nil, err
	}
	return &oss.PreSignedPutObjectUrlResp{
		PreSignedUrl: objUrl,
		ResourceUrl:  fmt.Sprintf("%s://%s.%s/%s", global.ServerConfig.OssConfig.AliYun.Scheme, req.BucketName, global.ServerConfig.OssConfig.AliYun.Endpoint, url.PathEscape(req.ObjectName)),
	}, nil
}

// PutObject 通过二进制文件 上传文件
func (ali *AliYun) PutObject(ctx context.Context, req *oss.PutObjectReq) (resp *oss.PutObjectResp, err error) {
	bucket, err := ali.Client.Bucket(req.BucketName)
	if err != nil {
		klog.CtxErrorf(ctx, "AliYun Oss Bucket error: %v", err)
		return nil, err
	}
	fileReader := bytes.NewReader(req.File)
	err = bucket.PutObject(req.ObjectName, fileReader, aliOss.ContentType("application/octet-stream"))
	if err != nil {
		klog.CtxErrorf(ctx, "AliYun Oss PutObject error: %v", err)
		return nil, err
	}
	return &oss.PutObjectResp{
		Url: fmt.Sprintf("%s.%s/%s", req.BucketName, global.ServerConfig.OssConfig.AliYun.Endpoint, req.ObjectName),
	}, nil
}
