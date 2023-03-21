// mq mq 服务
namespace go mq

struct Topic {
	uint32 created_at = 1;
	uint32 updated_at = 2;

	// @v: required
	string name = 3;
	repeated Channel channel_list = 4;

	// @desc: pub相关参数
	uint32 max_msg_size = 10;

	// @desc: topic 用哪组服务
	string node_group = 11;

	SubConfig sub_config = 20;

	string update_by = 21;

	// @desc: =true时，topic是不需消费的，通常是用于发布者自己不需要此消息，广播事件给需求方
	bool skip_consume_topic = 22;
}

struct AddTopicReq {
    1: Topic topic;
}
struct AddTopicResp {
    1: Topic topic;
}

service MqService {
    // 新增 topic
    AddTopicResp AddTopic(1: AddTopicReq req)
}


