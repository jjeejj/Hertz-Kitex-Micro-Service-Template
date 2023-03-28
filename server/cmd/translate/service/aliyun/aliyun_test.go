package aliyun

import (
	"context"
	"reflect"
	"testing"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/translator"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/log"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/translate/global"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/translate/initialize"
)

func init() {
	initialize.InitNacos()
	log.InitKLogger(consts.KlogFilePath, global.ServerConfig.LogLevel)
	initialize.InitAliYun()
}

func TestAliYun_DetectLanguage(t *testing.T) {
	type args struct {
		ctx context.Context
		req *translator.DetectLanguageReq
	}
	tests := []struct {
		name    string
		args    args
		want    *translator.DetectLanguageReq
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ali := &AliYun{}
			got, err := ali.DetectLanguage(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectLanguage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectLanguage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAliYun_Translate(t *testing.T) {
	type args struct {
		ctx context.Context
		req *translator.TranslateReq
	}
	tests := []struct {
		name    string
		args    args
		want    *translator.TranslateResp
		wantErr bool
	}{
		{
			name: "翻译中文到英文",
			args: args{
				ctx: context.Background(),
				req: &translator.TranslateReq{
					SourceLanguage: "zh",
					TargetLanguage: "en",
					SourceText:     "翻译中文到英文",
				},
			},
			want: &translator.TranslateResp{
				Result_: "Translate Chinese to English",
			},
		},
		{
			name: "翻译英文到中文",
			args: args{
				ctx: context.Background(),
				req: &translator.TranslateReq{
					SourceLanguage: "en",
					TargetLanguage: "zh",
					SourceText:     "Translate English to Chinese",
				},
			},
			want: &translator.TranslateResp{
				Result_: "将英语翻译成中文",
			},
		},
		{
			name: "自动识别语言到英文",
			args: args{
				ctx: context.Background(),
				req: &translator.TranslateReq{
					SourceLanguage: "auto",
					TargetLanguage: "en",
					SourceText:     "翻译中文到英文",
				},
			},
			want: &translator.TranslateResp{
				Result_: "Translate Chinese to English",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ali := &AliYun{}
			got, err := ali.Translate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
