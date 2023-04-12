package aliyun

import (
	"bufio"
	"context"
	"os"
	"reflect"
	"testing"

	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/initialize"
)

func init() {
	initialize.InitNacos()
	initialize.InitAliYun()
}

func TestAliYun_PreSignedPutObjectUrl(t *testing.T) {
	type fields struct {
		Client *aliOss.Client
	}
	type args struct {
		ctx context.Context
		req *oss.PreSignedPutObjectUrlReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *oss.PreSignedPutObjectUrlResp
		wantErr bool
	}{
		{
			name: "阿里云获取预签名url",
			fields: fields{
				Client: global.AliYunClient,
			},
			args: args{
				ctx: context.Background(),
				req: &oss.PreSignedPutObjectUrlReq{
					BucketName: "j-resource",
					ObjectName: "test.png",
					Type:       oss.OssPlatformType_ALI_YUN,
					Expiry:     1800,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ali := &AliYun{
				Client: tt.fields.Client,
			}
			got, err := ali.PreSignedPutObjectUrl(tt.args.ctx, tt.args.req)
			t.Logf("PreSignedPutObjectUrl: %v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("PreSignedPutObjectUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("PreSignedPutObjectUrl() got = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestAliYun_PutObject(t *testing.T) {
	fileByte, _ := os.Open("aliyun.go")
	defer fileByte.Close()
	stat, _ := fileByte.Stat()
	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, _ = bufio.NewReader(fileByte).Read(bs)
	type fields struct {
		Client *aliOss.Client
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
			fields: fields{Client: global.AliYunClient},
			args: args{
				ctx: context.Background(),
				req: &oss.PutObjectReq{
					BucketName: "j-resource",
					ObjectName: "aliyun.go",
					File:       bs,
					Type:       oss.OssPlatformType_ALI_YUN,
				},
			},
			wantResp: &oss.PutObjectResp{
				Url: "j-resource.oss-cn-shenzhen.aliyuncs.com/aliyun.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ali := &AliYun{
				Client: tt.fields.Client,
			}
			gotResp, err := ali.PutObject(tt.args.ctx, tt.args.req)
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
