package initialize

import (
	"flag"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/shared/consts"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/tool"
)

// InitFlag to init flag
func InitFlag() (string, int) {
	IP := flag.String(consts.IPFlagName, consts.IPFlagValue, consts.IPFlagUsage)
	Port := flag.Int(consts.PortFlagName, 0, consts.PortFlagUsage)
	// Parsing flags and if Port is 0 , then will automatically get an empty Port.
	flag.Parse()
	if *Port == 0 {
		*Port, _ = tool.GetFreePort()
	}
	klog.Info("ip: ", *IP)
	klog.Info("port: ", *Port)
	return *IP, *Port
}
