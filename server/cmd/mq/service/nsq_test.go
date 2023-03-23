package service

import (
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nsqio/go-nsq"
)

func TestSnqPublish(t *testing.T) {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		klog.Debug(err)
	}
	err = p.Publish("testTopic", []byte("nihao"))
	if err != nil {
		klog.Debug(err)
	}
	time.Sleep(time.Second * 1)
	p.Stop()

	// nsq.Consumer{}()
}
