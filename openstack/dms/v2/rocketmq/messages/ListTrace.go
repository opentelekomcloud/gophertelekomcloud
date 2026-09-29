package messages

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListTraceOpts struct {
	// Message ID.
	MsgID string `q:"msg_id,required"`
	// Number of records to query. Default: 10.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListTraceResponse is returned by ListTrace.
type ListTraceResponse struct {
	// Total number.
	Total int `json:"total"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
	// Message trace list.
	Trace []Trace `json:"trace"`
}

type Trace struct {
	// Whether the operation is successful.
	Success bool `json:"success"`
	// Trace type.
	//    Pub: The producer successfully sends messages.
	//    SubBefore: The consumer is ready to consume messages.
	//    SubAfter: The consumer finishes consuming messages.
	//    EndTransaction: Transactional messages are committed or rolled back.
	//    Receive: The service side receives messages.
	//    Ack: The consumer manually acknowledges consumption.
	TraceType string `json:"trace_type"`
	// Time.
	Timestamp int64 `json:"timestamp"`
	// Producer group or consumer group.
	GroupName string `json:"group_name"`
	// Duration.
	CostTime int64 `json:"cost_time"`
	// Request ID.
	RequestID string `json:"request_id"`
	// Consumption status.
	//    0: successful.
	//    1: timeout.
	//    2: exception.
	//    3: null.
	//    5: failed.
	ConsumeStatus int `json:"consume_status"`
	// Topic name.
	Topic string `json:"topic"`
	// Message ID.
	MsgID string `json:"msg_id"`
	// Offset message ID.
	OffsetMsgID string `json:"offset_msg_id"`
	// Message tag.
	Tags string `json:"tags"`
	// Message keys.
	Keys string `json:"keys"`
	// IP address of the host that stores the message.
	StoreHost string `json:"store_host"`
	// IP address of the host that generates the message.
	ClientHost string `json:"client_host"`
	// Number of retry times.
	RetryTimes int `json:"retry_times"`
	// Message body length.
	BodyLength int64 `json:"body_length"`
	// Message type.
	//    Normal_Msg: normal message.
	//    Trans_Msg_Half: half message.
	//    Trans_msg_Commit: delivered message.
	//    Delay_Msg: delayed message.
	//    Order_Msg: ordered message.
	MsgType string `json:"msg_type"`
	// Transaction status.
	//    COMMIT_MESSAGE
	//    ROLLBACK_MESSAGE
	//    UNKNOW
	TransactionState string `json:"transaction_state"`
	// Transaction ID.
	TransactionID string `json:"transaction_id"`
	// Whether the response is a transaction check response.
	FromTransactionCheck bool `json:"from_transaction_check"`
}

// ListTrace queries the trace of a message of a RocketMQ instance.
// Send GET to /v2/{project_id}/rocketmq/instances/{instance_id}/trace
func ListTrace(client *golangsdk.ServiceClient, instanceID string, opts ListTraceOpts) (*ListTraceResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "trace").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListTraceResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
