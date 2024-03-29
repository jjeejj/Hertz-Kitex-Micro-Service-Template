// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"

	_ "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/open_api/docs"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/open_api/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
}
