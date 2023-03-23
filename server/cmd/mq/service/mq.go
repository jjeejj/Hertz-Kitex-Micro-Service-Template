package service

import (
	"context"
	"reflect"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/mq"
)

// TopicMsg 主题定义.
type TopicMsg struct {
	Name        string        // 主题名称
	msgType     reflect.Type  // 消息类型
	channelList []*mq.Channel //  channel 列表
}

// NewTopicMsg 实例话 topic 相关调用的结构体.
func NewTopicMsg(topicName string, i interface{}) *TopicMsg {
	p := &TopicMsg{
		Name: topicName,
	}
	t := reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("invalid type, expect struct")
	}
	p.msgType = t

	return p
}

// Pub 发送消息.
func (t *TopicMsg) Pub(ctx context.Context, msg interface{}) {
	var buf []byte
	var err error
	buf, err = sonic.Marshal(msg)
	if err != nil {
		klog.Errorf("err:%v", err)

		return
	}
	klog.Infof("buf%v", buf)
}

// Process 处理消息.
func (t *TopicMsg) Process(ctx context.Context, msg interface{}) {
	var buf []byte
	var err error
	buf, err = sonic.Marshal(msg)
	if err != nil {
		klog.Errorf("err:%v", err)

		return
	}
	klog.Infof("buf%v", buf)
}
