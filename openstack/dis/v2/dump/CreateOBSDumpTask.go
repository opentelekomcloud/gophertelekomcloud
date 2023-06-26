package dump

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOBSDumpTaskOpts struct {
	// Name of the stream.
	// Maximum: 60
	StreamName string
	// Dump destination.
	// Possible values:
	// - OBS: Data is dumped to OBS.
	// Default: NOWHERE
	// Enumeration values:
	// OBS
	DestinationType string `json:"destination_type"`
	// Parameter list of OBS to which data in the DIS stream will be dumped.
	OBSDestinationDescriptor OBSDestinationDescriptorOpts `json:"obs_destination_descriptor,omitempty"`
}

func CreateOBSDumpTask(client *golangsdk.ServiceClient, opts CreateOBSDumpTaskOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/streams/{stream_name}/transfer-tasks
	_, err = client.Post(client.ServiceURL("streams", opts.StreamName, "transfer-tasks"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return err
}

type OBSDestinationDescriptorOpts struct {
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
	//	 Project: OBS. Select the Tenant Administrator role for the global service project.
	// If agencies have been created, you can obtain available agencies from the agency list by using the "Listing Agencies " API.
	// This parameter cannot be left blank and the parameter value cannot exceed 64 characters.
	// If there are dump tasks on the console, the system displays a message indicating that an agency will be automatically created.
	// The name of the automatically created agency is dis_admin_agency.
	// Maximum: 6
	AgencyName string `json:"agency_name"`
	// User-defined interval at which data is imported from the current DIS stream into OBS.
	// If no data is pushed to the DIS stream during the current interval, no dump file package will be generated.
	// Value range: 30-900
	// Default value: 300
	// Unit: second
	// Minimum: 30
	// Maximum: 900
	// Default: 300
	DeliverTimeInterval *int `json:"deliver_time_interval"`
	// Offset.
	// LATEST: Maximum offset indicating that the latest data will be extracted.
	// TRIM_HORIZON: Minimum offset indicating that the earliest data will be extracted.
	// Default value: LATEST
	// Default: LATEST
	// Enumeration values:
	// LATEST
	// TRIM_HORIZON
	ConsumerStrategy string `json:"consumer_strategy,omitempty"`
	// Directory to store files that will be dumped to OBS.
	// Different directory levels are separated by slashes (/) and cannot start with slashes.
	// The value can contain a maximum of 50 characters, including letters, digits, underscores (_), and slashes (/).
	// This parameter is left empty by default.
	// Maximum: 50
	FilePrefix string `json:"file_prefix,omitempty"`
	// Directory structure of the object file written into OBS.
	// The directory structure is in the format of yyyy/MM/dd/HH/mm (time at which the dump task was created).
	// - N/A: Leave this parameter empty, indicating that the date and time directory is not used.
	// - yyyy: year
	// - yyyy/MM: year/month
	// - yyyy/MM/dd: year/month/day
	// - yyyy/MM/dd/HH: year/month/day/hour
	// - yyyy/MM/dd/HH/mm: year/month/day/hour/minute
	// Example: in 2017/11/10/14/49, the directory structure is 2017 > 11 > 10 > 14 > 49. 2017 indicates the outermost folder.
	// Default value: empty.
	// Note:
	// After data is successfully dumped, the directory structure is obs_bucket_path/file_prefix/partition_format.
	// Enumeration values:
	// yyyy
	// yyyy/MM
	// yyyy/MM/dd
	// yyyy/MM/dd/HH
	// yyyy/MM/dd/HH/mm
	PartitionFormat string `json:"partition_format,omitempty"`
	// Name of the OBS bucket used to store data from the DIS stream.
	OBSBucketPath string `json:"obs_bucket_path,omitempty"`
	// Dump file format.
	// Possible values:
	// - Text (default)
	// - Parquet
	// - CarbonData
	// Note:
	// You can select Parquet or CarbonData only when Source Data Type is set to JSON an Dump Destination is set to OBS.
	// Default: text
	// Enumeration values:
	// text
	// parquet
	// carbon
	DestinationFileType string `json:"destination_file_type,omitempty"`
	// Dump time directory generated based on the timestamp of the source data and the configured partition_format.
	// Directory structure of the object file written into OBS.
	// The directory structure is in the format of yyyy/MM/dd/HH/mm.
	ProcessingSchema ProcessingSchema `json:"processing_schema,omitempty"`
	// Delimiter for the dump file which is used to separate the user data that is written into the dump file
	// Value range:
	// - Comma (,), which is the default value
	// - Semicolon (;)
	// - Vertical bar (|)
	// - Newline character (\n)
	// Default: \n
	RecordDelimiter string `json:"record_delimiter,omitempty"`
}

type ProcessingSchema struct {
	// Attribute name of the source data timestamp.
	TimestampName string `json:"timestamp_name"`
	// Type of the source data timestamp.
	// - String
	// - Timestamp: 13-bit timestamp of the long type
	TimestampType string `json:"timestamp_type"`
	// OBS directory generated based on the timestamp format.
	// This parameter is mandatory when the timestamp type of the source data is String.
	// Value range:
	// - yyyy/MM/dd HH:mm:ss
	// - MM/dd/yyyy HH:mm:ss
	// - dd/MM/yyyy HH:mm:ss
	// - yyyy-MM-dd HH:mm:ss
	// - MM-dd-yyyy HH:mm:ss
	// - dd-MM-yyyy HH:mm:ss
	// Enumeration values:
	// yyyy/MM/dd HH:mm:ss
	// MM/dd/yyyy HH:mm:ss
	// dd/MM/yyyy HH:mm:ss
	// yyyy-MM-dd HH:mm:ss
	// MM-dd-yyyy HH:mm:ss
	// dd-MM-yyyy HH:mm:ss
	TimestampFormat string `json:"timestamp_format,omitempty"`
}
