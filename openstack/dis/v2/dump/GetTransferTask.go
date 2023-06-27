package dump

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetTransferTaskOpts struct {
	// Name of the stream.
	StreamName string
	// Name of the dump task to be deleted.
	TaskName string
}

func GetTransferTask(client *golangsdk.ServiceClient, opts GetTransferTaskOpts) (*GetTransferTaskResponse, error) {
	// GET /v2/{project_id}/streams/{stream_name}/transfer-tasks/{task_name}
	raw, err := client.Get(client.ServiceURL("streams", opts.StreamName, "transfer-tasks", opts.TaskName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetTransferTaskResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetTransferTaskResponse struct {
	// Name of the stream to which the dump task belongs.
	StreamName string `json:"stream_name,omitempty"`
	// Name of the dump task.
	TaskName string `json:"task_name,omitempty"`
	// Id of the dump task
	TaskId string `json:"task_id,omitempty"`
	// Dump task status.
	// Possible values:
	// - ERROR: An error occurs.
	// - STARTING: The dump task is being started.
	// - PAUSED: The dump task has been stopped.
	// - RUNNING: The dump task is running.
	// - DELETE: The dump task has been deleted.
	// - ABNORMAL: The dump task is abnormal.
	// Enumeration values:
	//  ERROR
	//  STARTING
	//  PAUSED
	//  RUNNING
	//  DELETE
	//  ABNORMAL
	State string `json:"state,omitempty"`
	// Dump destination.
	// Possible values:
	// - OBS: Data is dumped to OBS.
	// - MRS: Data is dumped to MRS.
	// - DLI: Data is dumped to DLI.
	// - CLOUDTABLE: Data is dumped to CloudTable.
	// - DWS: Data is dumped to DWS.
	// Enumeration values:
	//  OBS
	//  MRS
	//  DLI
	//  CLOUDTABLE
	//  DWS
	DestinationType string `json:"destination_type,omitempty"`
	// Time when the dump task is created.
	CreatedAt *int64 `json:"create_time,omitempty"`
	// Latest dump time of the dump task.
	LastTransferTimestamp *int64 `json:"last_transfer_timestamp,omitempty"`
	// List of partition dump details.
	Partitions []PartitionResult `json:"partitions,omitempty"`
	// Parameter list of OBS to which data in the DIS stream will be dumped.
	OBSDestinationDescription OBSDestinationDescriptorOpts `json:"obs_destination_description,omitempty"`
	// Parameter list of the DWS to which data in the DIS stream will be dumped.
	DWSDestinationDescription DWSDestinationDescriptorOpts `json:"dws_destination_description,omitempty"`
	// Parameter list of the MRS to which data in the DIS stream will be dumped.
	MRSDestinationDescription MRSDestinationDescriptorOpts `json:"mrs_destination_description,omitempty"`
	// Parameter list of the DLI to which data in the DIS stream will be dumped.
	DLIDestinationDescription DLIDestinationDescriptorOpts `json:"dli_destination_description,omitempty"`
	// Parameter list of the CloudTable to which data in the DIS stream will be dumped.
	CloudTableDestinationDescription CloudTableDestinationDescriptorOpts `json:"cloud_table_destination_description,omitempty"`
}

type PartitionResult struct {
	// Current status of the partition.
	// Possible values:
	// - CREATING: The stream is being created.
	// - ACTIVE: The stream is available.
	// - DELETED: The stream is being deleted.
	// - EXPIRED: The stream has expired.
	// Enumeration values:
	//  CREATING
	//  ACTIVE
	//  DELETED
	//  EXPIRED
	Status string `json:"status,omitempty"`
	// Unique identifier of the partition.
	PartitionId string `json:"partition_id,omitempty"`
	// Possible value range of the hash key used by the partition.
	HashRange string `json:"hash_range,omitempty"`
	// Sequence number range of the partition.
	SequenceNumberRange string `json:"sequence_number_range,omitempty"`
	// Parent partition.
	ParentPartitions string `json:"parent_partitions,omitempty"`
}
