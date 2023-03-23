package service

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nsqio/go-nsq"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/mq/global"
)

type Nsq struct {
	publisher *nsq.Producer
	consumer  *nsq.Consumer
}

// Pub 发布消息.
func (n *Nsq) Pub(msg interface{}) {
	buf, err := sonic.Marshal(msg)
	if err != nil {
		klog.Errorf("err:%v", err)

		return
	}
	_ = n.publisher.Publish("", buf)
}

// Process 处理消息.
func (n *Nsq) Process(msg interface{}) {
	buf, err := sonic.Marshal(msg)
	if err != nil {
		klog.Errorf("err:%v", err, buf)

		return
	}
	_ = n.consumer.ConnectToNSQLookupds(global.ServerConfig.MqServerConfig.NSQLookupds)
	// n.consumer.AddHandler()
}
