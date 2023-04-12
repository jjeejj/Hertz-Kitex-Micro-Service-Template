package minio

import (
	"bufio"
	"context"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	minioGo "github.com/minio/minio-go/v7"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/initialize"
)

func init() {
	initialize.InitNacos()
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

func TestMinio_PutObject(t *testing.T) {
	fileByte, _ := os.Open("minio.go")
	defer fileByte.Close()
	stat, err := fileByte.Stat()
	if err != nil {
		t.Error(err)
		return
	}
	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(fileByte).Read(bs)
	if err != nil && err != io.EOF {
		t.Error(err)
		return
	}
	type fields struct {
		Client *minioGo.Client
	}
	type args struct {
		ctx context.Context
		req *oss.PutObjectReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *oss.PutObjectResp
		wantErr  bool
	}{
		{
			name:   "上传二进制",
			fields: fields{Client: global.MinioClient},
			args: args{
				ctx: context.Background(),
				req: &oss.PutObjectReq{
					BucketName: "test",
					ObjectName: "minio.go",
					File:       bs,
				},
			},
			wantResp: &oss.PutObjectResp{
				Url: "http://127.0.0.1:9000/test/minio.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minio := &Minio{
				Client: tt.fields.Client,
			}
			gotResp, err := minio.PutObject(tt.args.ctx, tt.args.req)
			t.Log("gotResp:", gotResp)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("PutObject() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
