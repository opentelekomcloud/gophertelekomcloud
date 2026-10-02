package diagnosis

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Report contains the details of an instance diagnosis report.
type Report struct {
	// Report ID.
	ReportID string `json:"report_id"`
	// Consumer group name.
	GroupName string `json:"group_name"`
	// Number of consumers.
	ConsumerNums int `json:"consumer_nums"`
	// Report status: diagnosing, failed, deleted, finished, normal or abnormal.
	Status string `json:"status"`
	// Generation time, Unix timestamp in milliseconds.
	CreatedAt int64 `json:"created_at"`
	// Number of abnormal items.
	AbnormalItemSum int `json:"abnormal_item_sum"`
	// Number of abnormal nodes.
	FaultedNodeSum int `json:"faulted_node_sum"`
	// Whether the consumer group is online.
	Online bool `json:"online"`
	// Number of stacked messages.
	MessageAccumulation int `json:"message_accumulation"`
	// Whether the subscription relationship is consistent.
	SubscriptionConsistency bool `json:"subscription_consistency"`
	// Whether there are duplicate client IDs.
	DuplicateClientID bool `json:"duplicate_client_id"`
	// Whether there are inconsistent consumption types.
	DifferentConsumerType bool `json:"different_consumer_type"`
	// Subscriber list.
	Subscriptions []Subscription `json:"subscriptions"`
	// Diagnosis node report list.
	NodeReports []NodeReport `json:"diagnosis_node_report_list"`
}

type Subscription struct {
	// Topic name.
	TopicName string `json:"topic_name"`
	// Consumer tag list.
	ConsumersInTags []ConsumersInTag `json:"consumers_in_tags"`
}

type ConsumersInTag struct {
	// Consumer list.
	Consumers []string `json:"consumers"`
	// Tag name.
	TagName string `json:"tag_name"`
}

type NodeReport struct {
	// Node ID.
	NodeID string `json:"node_id"`
	// Whether the node is faulty.
	IsFaulted bool `json:"is_faulted"`
	// Total number of exceptions.
	AbnormalItemSum int `json:"abnormal_item_sum"`
	// Number of stacked messages.
	MessageAccumulation int `json:"message_accumulation"`
	// Whether a deadlock occurs.
	DeadLock bool `json:"dead_lock"`
	// Deadlock thread.
	DeadlockThread string `json:"deadlock_thread"`
	// Stack ID, used to query stack information with GetStack.
	StackID string `json:"stack_id"`
	// Whether pop consumption is used.
	IsPop bool `json:"is_pop"`
	// Consumption type.
	ConsumeType string `json:"consume_type"`
}

// Get an instance diagnosis report.
// Send GET to /v2/{project_id}/rocketmq/diagnosis/{report_id}
func Get(client *golangsdk.ServiceClient, reportID string) (*Report, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("rocketmq", "diagnosis", reportID).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Report
	err = extract.Into(raw.Body, &res)
	return &res, err
}
