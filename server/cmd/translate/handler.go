package main

import (
	"context"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/translate"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/translate/service"
)

// TranslatorServiceImpl implements the last service interface defined in the IDL.
type TranslatorServiceImpl struct{}

// DetectLanguage implements the TranslatorServiceImpl interface.
func (s *TranslatorServiceImpl) DetectLanguage(ctx context.Context, req *translate.DetectLanguageReq) (resp *translate.DetectLanguageResp, err error) {
	translateClient, err := service.GetTranslateImpl(req.Type)
	if err != nil {
		return nil, err
	}
	resp, err = translateClient.DetectLanguage(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

// Translate implements the TranslatorServiceImpl interface.
func (s *TranslatorServiceImpl) Translate(ctx context.Context, req *translate.TranslateReq) (resp *translate.TranslateResp, err error) {
	translateClient, err := service.GetTranslateImpl(req.Type)
	if err != nil {
		return nil, err
	}
	resp, err = translateClient.Translate(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
