// Code generated by Kitex v0.4.4. DO NOT EDIT.

package authservice

import (
	"context"
	auth "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/api/kitex_gen/auth"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

var authServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Login":        kitex.NewMethodInfo(loginHandler, newAuthServiceLoginArgs, newAuthServiceLoginResult, false),
		"UploadAvatar": kitex.NewMethodInfo(uploadAvatarHandler, newAuthServiceUploadAvatarArgs, newAuthServiceUploadAvatarResult, false),
		"UpdateUser":   kitex.NewMethodInfo(updateUserHandler, newAuthServiceUpdateUserArgs, newAuthServiceUpdateUserResult, false),
		"GetUser":      kitex.NewMethodInfo(getUserHandler, newAuthServiceGetUserArgs, newAuthServiceGetUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "auth",
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

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceLoginArgs)
	realResult := result.(*auth.AuthServiceLoginResult)
	success, err := handler.(auth.AuthService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceLoginArgs() interface{} {
	return auth.NewAuthServiceLoginArgs()
}

func newAuthServiceLoginResult() interface{} {
	return auth.NewAuthServiceLoginResult()
}

func uploadAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceUploadAvatarArgs)
	realResult := result.(*auth.AuthServiceUploadAvatarResult)
	success, err := handler.(auth.AuthService).UploadAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceUploadAvatarArgs() interface{} {
	return auth.NewAuthServiceUploadAvatarArgs()
}

func newAuthServiceUploadAvatarResult() interface{} {
	return auth.NewAuthServiceUploadAvatarResult()
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceUpdateUserArgs)
	realResult := result.(*auth.AuthServiceUpdateUserResult)
	success, err := handler.(auth.AuthService).UpdateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceUpdateUserArgs() interface{} {
	return auth.NewAuthServiceUpdateUserArgs()
}

func newAuthServiceUpdateUserResult() interface{} {
	return auth.NewAuthServiceUpdateUserResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceGetUserArgs)
	realResult := result.(*auth.AuthServiceGetUserResult)
	success, err := handler.(auth.AuthService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceGetUserArgs() interface{} {
	return auth.NewAuthServiceGetUserArgs()
}

func newAuthServiceGetUserResult() interface{} {
	return auth.NewAuthServiceGetUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *auth.LoginRequest) (r *auth.LoginResponse, err error) {
	var _args auth.AuthServiceLoginArgs
	_args.Req = req
	var _result auth.AuthServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadAvatar(ctx context.Context, req *auth.UploadAvatarRequset) (r *auth.UploadAvatarResponse, err error) {
	var _args auth.AuthServiceUploadAvatarArgs
	_args.Req = req
	var _result auth.AuthServiceUploadAvatarResult
	if err = p.c.Call(ctx, "UploadAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, req *auth.UpdateUserRequest) (r *auth.UpdateUserResponse, err error) {
	var _args auth.AuthServiceUpdateUserArgs
	_args.Req = req
	var _result auth.AuthServiceUpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, req *auth.GetUserRequest) (r *auth.UserInfo, err error) {
	var _args auth.AuthServiceGetUserArgs
	_args.Req = req
	var _result auth.AuthServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
