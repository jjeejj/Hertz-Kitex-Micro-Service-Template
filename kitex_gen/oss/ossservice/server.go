// Code generated by Kitex v0.4.4. DO NOT EDIT.
package ossservice

import (
	server "github.com/cloudwego/kitex/server"
	oss "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler oss.OssService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}