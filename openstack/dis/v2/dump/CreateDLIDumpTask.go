package dump

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateDLIDumpTaskOpts struct {
	// Name of the stream.
	// Maximum: 60
	StreamName string
	// Dump destination.
	// Possible values:
	// - OBS: Data is dumped to OBS.
	// - MRS: Data is dumped to MRS.
	// - DLI: Data is dumped to DLI.
	// - CLOUDTABLE: Data is dumped to CloudTable.
	// - DWS: Data is dumped to DWS.
	// Default: NOWHERE
	// Enumeration values:
	// DLI
	DestinationType string `json:"destination_type"`
	// Parameter list of the DLI to which data in the DIS stream will be dumped.
	DLIDestinationDescriptor DLIDestinationDescriptorOpts `json:"mrs_destination_descriptor,omitempty"`
}

func CreateDLIDumpTask(client *golangsdk.ServiceClient, opts CreateDLIDumpTaskOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/streams/{stream_name}/transfer-tasks
	_, err = client.Post(client.ServiceURL("streams", opts.StreamName, "transfer-tasks"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}

type DLIDestinationDescriptorOpts struct {
	// Name of the dump task.
	// The task name consists of letters, digits, hyphens (-), and underscores (_).
	// It must be a string of 1 to 64 characters.
	TaskName string `json:"task_name"`
	// Name of the agency created on IAM.
	// DIS uses an agency to access your specified resources.
	// The parameters for creating an agency are as follows:
	// - Agency Type: Cloud service
	// - Cloud Service: DIS
	// - Validity Period: unlimited
	// - Scope: Global service,
	// 	 Project: OBS. Select the Tenant Administrator role for the global service project.
	// If agencies have been created, you can obtain available agencies from the agency list by using the "Listing Agencies " API.
	// This parameter cannot be left blank and the parameter value cannot exceed 64 characters.
	// If there are dump tasks on the console, the system displays a message indicating that an agency will be automatically created.
	// The name of the automatically created agency is dis_admin_agency.
	// Maximum: 64
	AgencyName string `json:"agency_name"`
	// User-defined interval at which data is imported from the current DIS stream into OBS.
	// If no data is pushed to the DIS stream during the current interval, no dump file package will be generated.
	// Value range: 30-900
	// Default value: 300
	// Unit: second
	// Minimum: 30
	// Maximum: 900
	// Default: 300
	DeliverTimeInterval string `json:"deliver_time_interval"`
	// Offset.
	// LATEST: Maximum offset indicating that the latest data will be extracted.
	// TRIM_HORIZON: Minimum offset indicating that the earliest data will be extracted.
	// Default value: LATEST
	// Default: LATEST
	// Enumeration values:
	// LATEST
	// TRIM_HORIZON
	ConsumerStrategy string `json:"consumer_strategy,omitempty"`
	// Name of the DLI database to which data in the DIS stream will be dumped.
	DLIDatabaseName string `json:"dli_database_name"`
	// Name of the DLI table to which data in the DIS stream will be dumped.
	// Note:
	// Only tables whose data location is DLI are supported, and you must have the permission to insert data into the tables.
	DLITableName string `json:"dli_table_name"`
	// Name of the OBS bucket used to temporarily store data in the DIS stream.
	OBSBucketPath string `json:"obs_bucket_path"`
	// Self-defined directory created in the OBS bucket and used to temporarily store data in the DIS stream.
	// Directory levels are separated by slashes (/) and cannot start with slashes.
	// The value can contain a maximum of 50 characters, including letters, digits, underscores (_), and slashes (/).
	// This parameter is left empty by default.
	FilePrefix string `json:"file_prefix,omitempty"`
	// Time duration for DIS to retry if data fails to be dumped to DLI.
	// If the retry time exceeds the value of this parameter, the data that fails to be dumped is backed up to the OBS bucket_path/file_prefix/mrs_error directory
	// Value range: 0-7,200
	// Unit: second
	// Default value: 1,800
	// If this parameter is set to 0, DIS does not retry when the dump fails.
	RetryDuration string `json:"retry_duration,omitempty"`
}
