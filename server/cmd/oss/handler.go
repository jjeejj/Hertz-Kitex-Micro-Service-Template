package main

import (
	"context"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/service"
)

// OssServiceImpl implements the last service interface defined in the IDL.
type OssServiceImpl struct{}

// PreSignedPutObjectUrl implements the OssServiceImpl interface.
func (s *OssServiceImpl) PreSignedPutObjectUrl(ctx context.Context, req *oss.PreSignedPutObjectUrlReq) (resp *oss.PreSignedPutObjectUrlResp, err error) {
	ossClient, err := service.GetOssImpl(req.Type)
	if err != nil {
		return nil, err
	}
	resp, err = ossClient.PreSignedPutObjectUrl(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

// PutObject implements the OssServiceImpl interface.
func (s *OssServiceImpl) PutObject(ctx context.Context, req *oss.PutObjectReq) (resp *oss.PutObjectResp, err error) {
	ossClient, err := service.GetOssImpl(req.Type)
	if err != nil {
		return nil, err
	}
	resp, err = ossClient.PutObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
