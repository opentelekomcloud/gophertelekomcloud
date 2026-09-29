package topics

// Topic contains the details of a RocketMQ topic.
type Topic struct {
	// Topic name.
	Name string `json:"name"`
	// Total number of read queues.
	TotalReadQueueNum int `json:"total_read_queue_num"`
	// Total number of write queues.
	TotalWriteQueueNum int `json:"total_write_queue_num"`
	// Permission.
	//    sub: subscribe permissions.
	//    pub: publish permissions.
	//    all: subscribe and publish permissions.
	Permission string `json:"permission"`
	// Associated brokers.
	Brokers []Broker `json:"brokers"`
	// Message type. Available only for RocketMQ 5.x instances.
	//    NORMAL: normal messages.
	//    FIFO: ordered messages.
	//    DELAY: scheduled messages.
	//    TRANSACTION: transactional messages.
	MessageType string `json:"message_type"`
	// Topic description.
	Description string `json:"topic_desc"`
	// Creation time, as the offset in milliseconds
	// from 1970-01-01 00:00:00 UTC to the specified time.
	CreatedAt int64 `json:"created_at"`
}

type Broker struct {
	// Broker name.
	BrokerName string `json:"broker_name"`
	// Number of read queues.
	ReadQueueNum int `json:"read_queue_num"`
	// Number of write queues.
	WriteQueueNum int `json:"write_queue_num"`
}
