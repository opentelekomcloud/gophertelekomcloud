package script

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type Script struct {
	// Workspace ID.
	Workspace string `json:"-"`
	// Name is a Script name. The name contains a maximum of 128 characters, including only letters, numbers, hyphens (-), and periods (.). The script name must be unique.
	Name string `json:"name" required:"true"`
	// Type is a script type. Can be: FlinkSQL, DLISQL, SparkSQL, HiveSQL, DWSSQL, RDSSQL, Shell, PRESTO, ClickHouseSQL, HetuEngineSQL, PYTHON, ImpalaSQL, Spark Python
	Type string `json:"type" required:"true"`
	// Content maximum of 4 MB is supported.
	Content string `json:"content,omitempty"`
	// Directory for storing the script. Access the DataArts Studio console and choose Data Development. The default directory is the root directory.
	Directory string `json:"directory,omitempty"`
	// ConnectionName is a name of the connection associated with the script.
	// This parameter is mandatory when type is set to DLISQL, SparkSQL, HiveSQL, DWSSQL, Shell, PRESTO, ClickHouseSQL, HetuEngineSQL, RDSSQL, ImpalaSQL, Spark Python, or PYTHON.
	// To obtain the existing connections, refer to the instructions in Querying a Connection List (to Be Taken Offline). By default, this parameter is left blank.
	ConnectionName string `json:"connectionName" required:"true"`
	// Database associated with an SQL statement. This parameter is available only when type is set to DLISQL, SparkSQL, HiveSQL, DWSSQL, PRESTO, RDSSQL, ClickHouseSQL, ImpalaSQL, or HetuEngineSQL.
	// If type is set to DLI SQL, obtain database information by calling the API for querying all databases in the Data Lake Insight API Reference.
	// This parameter is mandatory when the script is not of any type listed.
	Database string `json:"database,omitempty"`
	// Queue name of the DLI resource. This parameter is available only when type is set to DLISQL. You can obtain the queue information by calling the API for "Querying All Queues" in the Data Lake Insight API Reference. By default, this parameter is left blank.
	QueueName string `json:"queueName,omitempty"`
	// Configuration defined by a user for the job. This parameter is available only when type is set to DLISQL. For details about the supported configuration items, see conf parameter description in the "Submitting a SQL Job" section of the Data Lake Insight API Reference. By default, this parameter is left blank.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Description contains a maximum of 255 characters.
	Description string `json:"description,omitempty"`
	// This parameter is required if the review function is enabled. It indicates the target status of the script. The value can be SAVED, SUBMITTED, or PRODUCTION.
	TargetStatus string `json:"targetStatus,omitempty"`
	// Approvers is a script approver. This parameter is required if the review function is enabled.
	Approvers []*JobApprover `json:"approvers,omitempty"`
}

type JobApprover struct {
	ApproverName string `json:"approverName" required:"true"`
}

// Create is used to create a script. Currently, the following script types are supported: DLI SQL, Flink SQL, RDS SQL, Spark SQL, Hive SQL, DWS SQL, Shell, Presto SQL, ClickHouse SQL, HetuEngine SQL, Python, Spark Python, and Impala SQL.
// Send request POST /v1/{project_id}/scripts
func Create(client *golangsdk.ServiceClient, opts Script) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{HeaderContentType: ApplicationJson},
		OkCodes:     []int{204},
	}

	if opts.Workspace != "" {
		reqOpts.MoreHeaders[HeaderWorkspace] = opts.Workspace
	}

	_, err = client.Post(client.ServiceURL(scriptsEndpoint), b, nil, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
