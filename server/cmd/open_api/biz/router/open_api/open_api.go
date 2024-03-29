// Code generated by hertz generator. DO NOT EDIT.

package OpenApi

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	open_api "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/open_api/biz/handler/open_api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/token", append(_gettokenMw(), open_api.GetToken)...)
}
