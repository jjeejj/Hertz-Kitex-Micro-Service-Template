package initialize

import (
	mysql2 "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/pkg/mysql"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/global"
)

// InitDB to init database
func InitDB() {
	// global mode
	var err error
	global.DB, err = mysql2.InitDB(global.ServerConfig.MysqlInfo)
	if err != nil {
		klog.Fatalf("init mysql failed: %s", err.Error())
	}
}
