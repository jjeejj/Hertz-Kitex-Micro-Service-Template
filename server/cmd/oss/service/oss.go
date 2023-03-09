package service

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/service/minio"
)

type Oss interface {
	PreSignedPutObjectUrl(ctx context.Context, req *oss.PreSignedPutObjectUrlReq) (*oss.PreSignedPutObjectUrlResponse, error)
}

// GetOssImpl 获取 对应的 oss 的实现
func GetOssImpl(ossType oss.OssType) (Oss, error) {
	var ossClient Oss
	switch ossType {
	case oss.OssType_MINIO:
		ossClient = &minio.Minio{
			Client: global.MinioClient,
		}
	case oss.OssType_ALI_YUN:
	default:
		klog.Warnf("not support oss type: %v", ossType)
		return nil, errors.New("not support oss type")
	}
	return ossClient, nil
}
