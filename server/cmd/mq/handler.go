package main

import (
	"context"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/mq"
)

// MqServiceImpl implements the last service interface defined in the IDL.
type MqServiceImpl struct{}

// AddChannel implements the MqServiceImpl interface.
func (s *MqServiceImpl) AddChannel(ctx context.Context, req *mq.AddChannelReq) (resp *mq.AddChannelResp, err error) {
	// TODO: Your code here...
	return
}
