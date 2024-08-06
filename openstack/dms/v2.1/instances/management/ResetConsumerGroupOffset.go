package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const (
	resetMessageOffsetPath = "reset-message-offset"
)

type ResetMessageOffsetOpts struct {
	// Topic name.
	Topic string `json:"topic"`
	// Partition number. The default value is -1, indicating that all partitions are reset.
	Partition string `json:"partition" required:"true"`
	// Resetting consumer group offset to the specified position.
	//    If this position is earlier than the current earliest offset, the offset will be reset to the earliest offset.
	//    If this offset is later than the current largest offset, the offset will be reset to the latest offset.
	// Either message_offset or timestamp must be specified.
	MessageOffset int64 `json:"message_offset"`
	// Specified time that the offset is to be reset to. The value is a Unix timestamp, in millisecond.
	//    If this time is earlier than the current earliest timestamp, the offset will be reset to the earliest timestamp.
	//    If this time is later than the current largest timestamp, the offset will be reset to the latest timestamp.
	// Either message_offset or timestamp must be specified.
	Timestamp int64 `json:"timestamp" required:"true"`
}

// ResetconsumerGroupOffset Kafka instances do not support resetting the consumer offset online. Before resetting, stop the client for which the offset is to be reset.After a client is stopped, the server considers the client offline only after the time period specified in ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG (1000 ms by default).
// Send POST /v2/{project_id}/instances/{instance_id}/management/groups/{group}/reset-message-offset
func ResetconsumerGroupOffset(client *golangsdk.ServiceClient, instanceId, groupId string, opts ResetMessageOffsetOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL(instances.ResourcePath, instanceId, managementPath, groupPath, groupId, resetMessageOffsetPath), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
