// mq mq 服务
namespace go mq

// SubConfig 订阅的配置
struct SubConfig {
	1: string service_name; // 请求的服务
	2: string service_path; // 请求服务的路径
}
// Channel 主题的消费 channel
struct Channel {
	1: string name;
	2: SubConfig sub_config;
}

struct AddChannelReq {
    1: string topic_name;
    2: Channel channel;
}
struct AddChannelResp {
    1: Channel channel;
}

service MqService {
    // 新增 Channel
    AddChannelResp AddChannel(1: AddChannelReq req);
}


