package dump

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateDWSDumpTaskOpts struct {
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
	// DWS
	DestinationType string `json:"destination_type"`
	// Parameter list of the DWS to which data in the DIS stream will be dumped.
	DWSDestinationDescriptor DWSDestinationDescriptorOpts `json:"dws_destination_descriptor,omitempty"`
}

func CreateDWSDumpTask(client *golangsdk.ServiceClient, opts CreateDWSDumpTaskOpts) error {
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

type DWSDestinationDescriptorOpts struct {
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
	// Name of the DWS cluster that stores the data in the stream.
	DWSClusterName string `json:"dws_cluster_name"`
	// ID of the DWS cluster to which will be dumped.
	DWSClusterID string `json:"dws_cluster_id"`
	// Name of the DWS database that stores the data in the stream.
	DWSDatabaseName string `json:"dws_database_name"`
	// Schema of the DWS database to which data will be dumped.
	DWSSchema string `json:"dws_schema"`
	// Name of the DWS table that stores the data in the stream.
	DWSTableName string `json:"dws_table_name"`
	// Delimiter used to separate the columns in the DWS tables.
	// The value can be a comma (,), semicolon (;), or vertical bar (|).
	DWSDelimiter string `json:"dws_delimiter"`
	// Username of the DWS database to which data will be dumped.
	UserName string `json:"user_name"`
	// Password of the DWS database to which data will be dumped.
	UserPassword string `json:"user_password"`
	// Key created in Key Management Service (KMS) and used to encrypt the password of the DWS database.
	KMSUserKeyName string `json:"kms_user_key_name"`
	// ID of the key created in KMS and used to encrypt the password of the DWS database.
	KMSUserKeyID string `json:"kms_user_key_id"`
	// Name of the OBS bucket used to temporarily store data in the DIS stream.
	OBSBucketPath string `json:"obs_bucket_path"`
	// User-defined directory created in the OBS bucket and used to temporarily store data in the DIS stream.
	// Directory levels are separated by slashes (/) and cannot start with slashes.
	// The value can contain a maximum of 50 characters, including letters, digits, underscores (_), and slashes (/).
	// This parameter is left empty by default.
	FilePrefix string `json:"file_prefix,omitempty"`
	// Duration when you can constantly retry dumping data to DWS after the dump fails.
	// If the dump time exceeds the value of this parameter, the data that fails to be dumped to DWS will be backed up to the OBS bucket_path/file_prefix/dws_error directory.
	// Value range: 0-7,200
	// Unit: second
	// Default value: 1,800
	RetryDuration string `json:"retry_duration,omitempty"`
	// Column to be dumped to the DWS table.
	// If the value is null or empty, all columns are dumped by default.
	// For example, c1,c2 indicates that columns c1 and c2 in the schema are dumped to DWS.
	// This parameter is left blank by default.
	DWSTableColumns string `json:"dws_table_columns,omitempty"`
	// DWS fault tolerance option (used to specify various parameters of foreign table data).
	Options Options `json:"options,omitempty"`
}

type Options struct {
	// Specifies whether to set the field to Null or enable an error message to be displayed in the error table when the last field in a row of the data source file is missing during database import.
	// Value range:
	// - true/on
	// - false/off
	// Default value: false/off
	// Enumeration values:
	// true/on
	// false/off
	FillMissingFields string `json:"fill_missing_fields,omitempty"`
	// Specifies whether to ignore excessive columns when the number of columns in a source data file exceeds that defined in the foreign table.
	// This parameter is used only during data import.
	// Value range:
	// - true/on
	// - false/off
	// Default value: false/off
	// Enumeration values:
	// true/on
	// false/off
	IgnoreExtraData string `json:"ignore_extra_data,omitempty"`
	// Specifies whether to tolerate invalid characters during data import.
	// Specifies whether to convert invalid characters based on the conversion rule and import them to the database, or to report an error and stop the import.
	// Value range:
	// - true/on
	// - false/off
	// Default value: false/off
	// Enumeration values:
	// true/on
	// false/off
	CompatibleIllegalChars string `json:"compatible_illegal_chars,omitempty"`
	// Maximum number of data format errors allowed during the data import.
	// If the number of data format errors does not reach the maximum, the data import is successful.
	// Value range:
	// - integer
	// - unlimited
	// Default value: 0,
	// indicating that error information is returned immediately
	RejectLimit string `json:"reject_limit,omitempty"`
	// Name of the error table that records data format errors.
	// After the parallel import is complete, you can query the error information table to obtain the detailed error information.
	ErrorTableName string `json:"error_table_name,omitempty"`
}
