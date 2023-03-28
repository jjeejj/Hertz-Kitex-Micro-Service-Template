// Code generated by Kitex v0.4.4. DO NOT EDIT.

package translateservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	translate "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/translate"
)

func serviceInfo() *kitex.ServiceInfo {
	return translateServiceServiceInfo
}

var translateServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "TranslateService"
	handlerType := (*translate.TranslateService)(nil)
	methods := map[string]kitex.MethodInfo{
		"DetectLanguage": kitex.NewMethodInfo(detectLanguageHandler, newTranslateServiceDetectLanguageArgs, newTranslateServiceDetectLanguageResult, false),
		"Translate":      kitex.NewMethodInfo(translateHandler, newTranslateServiceTranslateArgs, newTranslateServiceTranslateResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "translate",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func detectLanguageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*translate.TranslateServiceDetectLanguageArgs)
	realResult := result.(*translate.TranslateServiceDetectLanguageResult)
	success, err := handler.(translate.TranslateService).DetectLanguage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTranslateServiceDetectLanguageArgs() interface{} {
	return translate.NewTranslateServiceDetectLanguageArgs()
}

func newTranslateServiceDetectLanguageResult() interface{} {
	return translate.NewTranslateServiceDetectLanguageResult()
}

func translateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*translate.TranslateServiceTranslateArgs)
	realResult := result.(*translate.TranslateServiceTranslateResult)
	success, err := handler.(translate.TranslateService).Translate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTranslateServiceTranslateArgs() interface{} {
	return translate.NewTranslateServiceTranslateArgs()
}

func newTranslateServiceTranslateResult() interface{} {
	return translate.NewTranslateServiceTranslateResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) DetectLanguage(ctx context.Context, req *translate.DetectLanguageReq) (r *translate.DetectLanguageResp, err error) {
	var _args translate.TranslateServiceDetectLanguageArgs
	_args.Req = req
	var _result translate.TranslateServiceDetectLanguageResult
	if err = p.c.Call(ctx, "DetectLanguage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Translate(ctx context.Context, req *translate.TranslateReq) (r *translate.TranslateResp, err error) {
	var _args translate.TranslateServiceTranslateArgs
	_args.Req = req
	var _result translate.TranslateServiceTranslateResult
	if err = p.c.Call(ctx, "Translate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
