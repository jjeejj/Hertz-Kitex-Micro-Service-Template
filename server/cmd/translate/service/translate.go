package service

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/translate"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/translate/global"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/translate/service/aliyun"
)

type Translate interface {
	// DetectLanguage 识别指定文本语言
	DetectLanguage(ctx context.Context, req *translate.DetectLanguageReq) (*translate.DetectLanguageResp, error)
	// Translate 翻译文本到指定语言
	Translate(ctx context.Context, req *translate.TranslateReq) (*translate.TranslateResp, error)
}

// GetTranslateImpl 获取 对应的  翻译 的实现.
func GetTranslateImpl(translateType translate.TranslatePlatformType) (Translate, error) {
	var translateClient Translate
	switch translateType {
	case translate.TranslatePlatformType_ALI_YUN:
		translateClient = &aliyun.AliYun{
			Client: global.AliYunClient,
		}
	default:
		klog.Warnf("not support translate type: %v", translateType)
		return nil, errors.New("not support translate type")
	}

	return translateClient, nil
}
