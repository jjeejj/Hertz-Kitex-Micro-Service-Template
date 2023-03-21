package global

import (
	"github.com/nsqio/go-nsq"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/mq/config"
)

var (
	ServerConfig config.ServerConfig
	SnqClient    *nsq.Consumer
)
