package topics

// CreateResponse is a struct that contains the create response
type CreateResponse struct {
	Name string `json:"name"`
}

// Topic includes the parameters of an topic
// REDO this part
type Topic struct {
	Name             string `json:"name"`
	Partition        int    `json:"partition"`
	Replication      int    `json:"replication"`
	RetentionTime    int    `json:"retention_time"`
	SyncReplication  bool   `json:"sync_replication"`
	SyncMessageFlush bool   `json:"sync_message_flush"`
	TopicType        int    `json:"topic_type"`
	PoliciesOnly     bool   `json:"policiesOnly"`
	ExternalConfigs  any    `json:"external_configs"`
}

// ListResponse is a struct that contains the list response
type ListResponse struct {
	Total            int     `json:"total"`
	Size             int     `json:"size"`
	RemainPartitions int     `json:"remain_partitions"`
	MaxPartitions    int     `json:"max_partitions"`
	Topics           []Topic `json:"topics"`
}

// TopicDetail includes the detail parameters of an topic
type TopicDetail struct {
	Name            string      `json:"topic"`
	Partitions      []Partition `json:"partitions"`
	GroupSubscribed []string    `json:"group_subscribed"`
}

// Partition represents the details of a partition
type Partition struct {
	Partition int       `json:"partition"`
	Replicas  []Replica `json:"replicas"`
	// Node ID
	Leader int `json:"leader"`
	// Log End Offset
	Leo int `json:"leo"`
	// High Watermark
	Hw int `json:"hw"`
	// Log Start Offset
	Lso int `json:"lso"`
	// time stamp
	UpdateTimestamp int64 `json:"last_update_timestamp"`
}

// Replica represents the details of a replica
type Replica struct {
	Broker int  `json:"broker"`
	Leader bool `json:"leader"`
	InSync bool `json:"in_sync"`
	Size   int  `json:"size"`
	Lag    int  `json:"lag"`
}

// DeleteResponse is a struct that contains the deletion response
type DeleteResponse struct {
	Name    string `json:"id"`
	Success bool   `json:"success"`
}
