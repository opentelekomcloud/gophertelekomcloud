package dump

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateCloudTableDumpTaskOpts struct {
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
	// MRS
	DestinationType string `json:"destination_type"`
	// Parameter list of the CloudTable to which data in the DIS stream will be dumped.
	CloudTableDestinationDescriptor CloudTableDestinationDescriptorOpts `json:"cloudtable_destination_descriptor,omitempty"`
}

func CreateCloudTableDumpTask(client *golangsdk.ServiceClient, opts CreateCloudTableDumpTaskOpts) error {
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

type CloudTableDestinationDescriptorOpts struct {
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
	// Name of the CloudTable cluster to which data will be dumped.
	// If you choose to dump data to OpenTSDB, OpenTSDB must be enabled for the cluster.
	CloudTableClusterName string `json:"cloudtable_cluster_name"`
	// ID of the CloudTable cluster to which data will be dumped.
	// If you choose to dump data to OpenTSDB, OpenTSDB must be enabled for the cluster.
	CloudTableClusterID string `json:"cloudtable_cluster_id"`
	// HBase table name of the CloudTable cluster to which data will be dumped.
	// The parameter is mandatory when data is dumped to the CloudTable HBase.
	CloudTableTableName string `json:"cloudtable_table_name,omitempty"`
	// Schema configuration of the CloudTable HBase data.
	// You can set either this parameter or opentsdb_schema, but this parameter is mandatory when data will be dumped to HBase.
	// After this parameter is set, the JSON data in the stream can be converted to another format and then be imported to the CloudTable HBase.
	CloudtableSchema CloudtableSchema `json:"cloudtable_schema,omitempty"`
	// Schema configuration of the CloudTable OpenTSDB data.
	// You can set either this parameter or opentsdb_schema, but this parameter is mandatory when data will be dumped to OpenTSDB.
	// After this parameter is set, the JSON data in the stream can be converted to another format and then be imported to the CloudTable OpenTSDB.
	OpenTSDBSchema OpenTSDBSchema `json:"opentsdb_schema,omitempty"`
	// Delimiter used to separate the user data that generates HBase row keys.
	// Value range: , . | ; \ - _ and ~
	// Default value: .
	CloudtableRowKeyDelimiter string `json:"cloudtable_row_key_delimiter,omitempty"`
	// Name of the OBS bucket used to back up data that failed to be dumped to CloudTable.
	OBSBackupBucketPath string `json:"obs_backup_bucket_path,omitempty"`
	// Self-defined directory created in the OBS bucket and used to back up data that failed to be dumped to CloudTable.
	// Directory levels are separated by slashes (/) and cannot start with slashes.
	// Value range: a string of letters, digits, and underscores (_)
	// The maximum length is 50 characters.
	// This parameter is left empty by default.
	BackupFilePrefix string `json:"backup_file_prefix,omitempty"`
	// Time duration for DIS to retry if data fails to be dumped to CloudTable.
	// If this threshold is exceeded, the data that fails to be dumped will be backed up to the OBS bucket/ backup_file_prefix / cloudtable_error or OBS bucket/ backup_file_prefix / opentsdb_error directory.
	// Value range: 0-7,200
	// Unit: second
	// Default value: 1,800
	RetryDuration string `json:"retry_duration,omitempty"`
}

type CloudtableSchema struct {
	// HBase rowkey schema used by the CloudTable cluster to convert JSON data into HBase rowkeys.
	// Value range: 1-64
	RowKey []RowKey `json:"row_key"`
	// HBase column schema used by the CloudTable cluster to convert JSON data into HBase columns.
	// Value range: 1 to 4,096
	Columns []Column `json:"columns"`
}

type RowKey struct {
	// JSON attribute name, which is used to generate HBase rowkeys for JSON data in the DIS stream.
	Value string `json:"value"`
	// JSON attribute type of JSON data in the DIS stream.
	// Value range:
	// - Bigint
	// - Double
	// - Boolean
	// - Timestamp
	// - String
	// - Decimal
	// Enumeration values:
	// Bigint
	// Double
	// Boolean
	// Timestamp
	// String
	// Decimal
	Type string `json:"type"`
}

type Column struct {
	// Name of the HBase column family to which data will be dumped.
	ColumnFamilyName string `json:"column_family_name"`
	// Name of the HBase column to which data will be dumped.
	// Value range: a string of 1 to 32 characters, consisting of only letters, digits, and underscores (_)
	ColumnName string `json:"column_name"`
	// JSON attribute name, which is used to generate HBase column values for JSON data in the DIS stream.
	Value string `json:"value"`
	// JSON attribute type of JSON data in the DIS stream.
	// Value range:
	// - Bigint
	// - Double
	// - Boolean
	// - Timestamp
	// - String
	// - Decimal
	// Enumeration values:
	// Bigint
	// Double
	// Boolean
	// Timestamp
	// String
	// Decimal
	Type string `json:"type"`
}

type OpenTSDBSchema struct {
	// Schema configuration of the OpenTSDB data metric in the CloudTable cluster.
	// After this parameter is set, the JSON data in the stream can be converted to the metric of the OpenTSDB data.
	Metric []OpenTSDBMetric `json:"metric"`
	// Schema configuration of the OpenTSDB data timestamp in the CloudTable cluster.
	// After this parameter is set, the JSON data in the stream can be converted to the timestamp of the OpenTSDB data.
	Timestamp OpenTSDBTimestamp `json:"timestamp"`
	// Schema configuration of the OpenTSDB data value in the CloudTable cluster.
	// After this parameter is set, the JSON data in the stream can be converted to the value of the OpenTSDB data.
	Value OpenTSDBValue `json:"value"`
	// Schema configuration of the OpenTSDB data tags in the CloudTable cluster.
	// After this parameter is set, the JSON data in the stream can be converted to the tags of the OpenTSDB data.
	Tags []OpenTSDBTags `json:"tags"`
}

type OpenTSDBMetric struct {
	// When type is set to Constant, the value of metric is the value of Value.
	// When value is set to String, the value of metric is the value of the JSON attribute of the user data in the stream.
	// Enumeration values:
	// Constant
	// String
	Type string `json:"type"`
	// Constant value or JSON attribute name of the user data in the stream.
	// This value is 1 to 32 characters long.
	// Only letters, digits, and periods (.) are allowed.
	Value string `json:"value"`
}

type OpenTSDBTimestamp struct {
	// When type is set to Timestamp, the value type of the JSON attribute of the user data in the stream is Timestamp, and the timestamp of OpenTSDB can be generated without converting the data format.
	// When type is set to String, the value type of the JSON attribute of the user data in the stream is Date, and the timestamp of OpenTSDB can be generated only after the data format is converted.
	Type string `json:"type"`
	// JSON attribute name of the user data in the stream.
	// Value range: a string of 1 to 32 characters, consisting of only letters, digits, and underscores (_)
	Value string `json:"value"`
	// This parameter is mandatory when type is set to String.
	// When the value type of the JSON attribute of the user data in the stream is Date, format is required to convert the data format to generate the timestamp of OpenTSDB.
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
	Format string `json:"format"`
}

type OpenTSDBValue struct {
	// Dump destination.
	// Possible values:
	// Value range:
	// - Bigint
	// - Double
	// - Boolean
	// - Timestamp
	// - String
	// - Decimal
	Type string `json:"type"`
	// Constant value or JSON attribute name of the user data in the stream.
	// Value range: a string of 1 to 32 characters, consisting of only letters, digits, and underscores (_)
	Value string `json:"value"`
}

type OpenTSDBTags struct {
	// Tag name of the OpenTSDB data that stores the data in the stream.
	// Value range: a string of 1 to 32 characters, consisting of only letters, digits, and underscores (_)
	Name string `json:"name"`
	// Type name of the JSON attribute of the user data in the stream.
	// Value range:
	// - Bigint
	// - Double
	// - Boolean
	// - Timestamp
	// - String
	// - Decimal
	Type string `json:"type"`
	// Constant value or JSON attribute name of the user data in the stream.
	// Value range: a string of 1 to 32 characters, consisting of only letters, digits, and underscores (_)
	Value string `json:"value"`
}
