package aliyun

import (
	"context"
	"errors"
	"net/http"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/translate"
	_const "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/translate/const"
)

// 使用通用翻译模型

type AliYun struct {
	Client *alimt.Client
}

func (ali *AliYun) DetectLanguage(ctx context.Context, req *translate.DetectLanguageReq) (*translate.DetectLanguageResp, error) {
	request := alimt.CreateGetDetectLanguageRequest()
	request.SourceText = req.SourceText
	respData, err := ali.Client.GetDetectLanguage(request)
	if err != nil {
		return nil, err
	}
	return &translate.DetectLanguageResp{
		DetectedLanguage: respData.DetectedLanguage,
	}, nil
}

func (ali *AliYun) Translate(ctx context.Context, req *translate.TranslateReq) (*translate.TranslateResp, error) {
	request := alimt.CreateTranslateRequest()
	request.SourceText = req.SourceText
	request.SourceLanguage = req.SourceLanguage
	request.Scene = _const.AliYun_Scene_General
	request.FormatType = _const.AliYun_FormatType_Text
	request.TargetLanguage = req.TargetLanguage
	request.Method = http.MethodPost
	respData, err := ali.Client.Translate(request)
	if err != nil {
		return nil, err
	}
	// 请求非 200
	if respData == nil || respData.Code != 200 {
		return nil, errors.New(respData.Message)
	}
	return &translate.TranslateResp{
		TranslateResult_: respData.Data.Translated,
	}, nil
}
