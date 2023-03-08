package main

import (
	"context"
	oss "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
)

// OssServiceImpl implements the last service interface defined in the IDL.
type OssServiceImpl struct{}

// PreSignedPutObjectUrl implements the OssServiceImpl interface.
func (s *OssServiceImpl) PreSignedPutObjectUrl(ctx context.Context, req *oss.PreSignedPutObjectUrlReq) (resp *oss.PreSignedPutObjectUrlResponse, err error) {
	// TODO: Your code here...
	return
}
