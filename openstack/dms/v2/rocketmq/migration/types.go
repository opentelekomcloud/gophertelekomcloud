package migration

// Task contains the details of a metadata migration task.
type Task struct {
	// Task ID.
	ID string `json:"id"`
	// Task name.
	Name string `json:"name"`
	// Start time of the task.
	StartDate string `json:"start_date"`
	// Task status.
	//    creating
	//    starting: the migration is in progress.
	//    failed
	//    finished
	Status string `json:"status"`
	// Migration type.
	//    rocketmq: from RocketMQ to RocketMQ.
	//    rabbitToRocket: from RabbitMQ to RocketMQ.
	Type string `json:"type"`
}

// TopicConfig is the metadata of a RocketMQ topic.
type TopicConfig struct {
	// Topic name.
	TopicName string `json:"topic_name,omitempty"`
	// Whether messages are ordered. Default: false.
	Order bool `json:"order,omitempty"`
	// Topic permissions. Default: 6.
	Perm int `json:"perm,omitempty"`
	// Number of read queues. Default: 16.
	ReadQueueNums int `json:"read_queue_nums,omitempty"`
	// Number of write queues. Default: 16.
	WriteQueueNums int `json:"write_queue_nums,omitempty"`
	// Topic filtering type: SINGLE_TAG or MULTI_TAG.
	TopicFilterType string `json:"topic_filter_type,omitempty"`
	// Topic system flag. Default: 0.
	TopicSysFlag int `json:"topic_sys_flag,omitempty"`
}

// SubscriptionGroup is the metadata of a RocketMQ consumer group.
type SubscriptionGroup struct {
	// Consumer group name.
	GroupName string `json:"group_name,omitempty"`
	// Whether to allow broadcast.
	ConsumeBroadcastEnable *bool `json:"consume_broadcast_enable,omitempty"`
	// Whether to enable consumption. Default: true.
	ConsumeEnable *bool `json:"consume_enable,omitempty"`
	// Whether to enable consumption from the earliest offset. Default: true.
	ConsumeFromMinEnable *bool `json:"consume_from_min_enable,omitempty"`
	// Whether to notify consumer ID changes. Default: true.
	NotifyConsumerIDsChangedEnable *bool `json:"notify_consumer_ids_changed_enable,omitempty"`
	// Maximum number of consumption retries. Default: 16.
	RetryMaxTimes int `json:"retry_max_times,omitempty"`
	// Number of retry queues. Default: 1.
	RetryQueueNums int `json:"retry_queue_nums,omitempty"`
	// ID of the broker selected for slow consumption. Default: 1.
	WhichBrokerWhenConsumeSlow int64 `json:"which_broker_when_consume_slow,omitempty"`
}

// Vhost is the metadata of a RabbitMQ virtual host.
type Vhost struct {
	// Virtual host name.
	Name string `json:"name,omitempty"`
}

// Queue is the metadata of a RabbitMQ queue.
type Queue struct {
	// Virtual host name.
	Vhost string `json:"vhost,omitempty"`
	// Queue name.
	Name string `json:"name,omitempty"`
	// Whether to enable data persistence.
	Durable bool `json:"durable"`
}

// Exchange is the metadata of a RabbitMQ exchange.
type Exchange struct {
	// Virtual host name.
	Vhost string `json:"vhost,omitempty"`
	// Exchange name.
	Name string `json:"name,omitempty"`
	// Exchange type: topic, direct, fanout or headers.
	Type string `json:"type,omitempty"`
	// Whether to enable data persistence.
	Durable bool `json:"durable"`
}

// Binding is the metadata of a RabbitMQ binding.
type Binding struct {
	// Virtual host name.
	Vhost string `json:"vhost,omitempty"`
	// Message source.
	Source string `json:"source,omitempty"`
	// Message target.
	Destination string `json:"destination,omitempty"`
	// Target type: exchange or queue.
	DestinationType string `json:"destination_type,omitempty"`
	// Routing key.
	RoutingKey string `json:"routing_key,omitempty"`
}
