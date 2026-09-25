package groups

// ConsumerGroup contains the details of a RocketMQ consumer group.
type ConsumerGroup struct {
	// Consumer group name.
	Name string `json:"name"`
	// Consumer group description.
	Description string `json:"group_desc"`
	// Whether consumption is allowed.
	Enabled bool `json:"enabled"`
	// Whether broadcast is enabled.
	Broadcast bool `json:"broadcast"`
	// List of associated brokers.
	Brokers []string `json:"brokers"`
	// Maximum number of retries.
	RetryMaxTime int `json:"retry_max_time"`
	// Creation time, as the offset in milliseconds
	// from 1970-01-01 00:00:00 UTC to the specified time.
	CreatedAt int64 `json:"created_at"`
	// Permission set.
	Permissions []string `json:"permissions"`
	// Whether orderly consumption is enabled.
	ConsumeOrderly bool `json:"consume_orderly"`
	// Whether the consumer group is online.
	GroupOnline bool `json:"group_online"`
}
