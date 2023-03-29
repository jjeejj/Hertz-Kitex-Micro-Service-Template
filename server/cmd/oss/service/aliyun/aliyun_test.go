package aliyun

import (
	"context"
	"reflect"
	"testing"

	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/oss/global"
)

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
					BucketName: "",
					ObjectName: "",
					Type:       oss.OssPlatformType_ALI_YUN,
					Expiry:     60,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ali := &AliYun{
				Client: tt.fields.Client,
			}
			got, err := ali.PreSignedPutObjectUrl(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PreSignedPutObjectUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreSignedPutObjectUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
