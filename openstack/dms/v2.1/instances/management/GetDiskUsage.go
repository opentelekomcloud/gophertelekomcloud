package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const (
	diskUsagePath = "diskusage"
)

type GetDiskUsageOpts struct {
	// Querying partitions by the used disk space. Options: 1 KB, 1 MB and 1 GB. Default value: 1 GB.
	MinSize int `q:"minSize,omitempty"`
	// Querying partitions by top disk usage.
	Top string `q:"top,omitempty"`
	// Querying partitions by the percentage of the used disk space.
	Percentage string `q:"percentage,omitempty"`
}

// GetDiskUsageStatusOfTopics is used to query the broker disk usage of topics.
// Send GET /v2/{project_id}/instances/{instance_id}/topics/diskusage
func GetDiskUsageStatusOfTopics(client *golangsdk.ServiceClient, instanceId string, opts GetDiskUsageOpts) (*GetDiskUsageStatusOfTopicsResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(instances.ResourcePath, instanceId, topicPath, diskUsagePath).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetDiskUsageStatusOfTopicsResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetDiskUsageStatusOfTopicsResp struct {
	// Broker list.
	Brokers []*DiskUsage `json:"broker_list"`
}

type DiskUsage struct {
	// Broker name.
	BrokerName string `json:"broker_name"`
	// Disk capacity.
	DatadiskSize string `json:"datadisk_size"`
	// Used disk space.
	DataDiskUsage string `json:"data_disk_use"`
	// Remaining disk space.
	DataDiskFree string `json:"data_disk_free"`
	// Message label.
	DataDiskUsagePercentage string `json:"data_disk_usage_percentage"`
	// Message label.
	Status string `json:"status"`
	// Disk usage list of the topics.
	TopicList []*DiskUsageTopic `json:"topic_list"`
}

type DiskUsageTopic struct {
	// Disk usage.
	Size string `json:"size"`
	// Topic name.
	TopicName string `json:"topic_name"`
	// Partition.
	TopicPartition string `json:"topic_partition"`
	// Percentage of used disk space.
	Percentage float64 `json:"percentage"`
}
