package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type Job struct {
	// Job name. The name contains a maximum of 128 characters, including only letters, numbers, hyphens (-), underscores (_), and periods (.). The job name must be unique.
	Name string `json:"name" required:"true"`
	// Nodes is a node definition.
	Nodes []Node `json:"nodes" required:"true"`
	// Schedule is a scheduling configuration.
	Schedule Schedule `json:"schedule" required:"true"`
	// Job parameter definition.
	Params []Param `json:"params,omitempty"`
	// Path of a job in the directory tree. If the directory of the path does not exist during job creation, a directory is automatically created in the root directory /, for example, /dir/a/.
	Directory string `json:"directory,omitempty"`
	// Job type.
	// REAL_TIME: real-time processing
	// BATCH: batch processing
	ProcessType string `json:"processType" required:"true"`
	// Job ID
	Id int64 `json:"id,omitempty"`
	// User who last updated the job
	LastUpdateUser string `json:"lastUpdateUser,omitempty"`
	// OBS path for storing job run logs
	LogPath string `json:"logPath,omitempty"`
	// Basic job information.
	BasicConfig BasicConfig `json:"basicConfig,omitempty"`
	// This parameter is required if the review function is enabled. It indicates the target status of the job. The value can be SAVED, SUBMITTED, or PRODUCTION.
	TargetStatus string `json:"targetStatus,omitempty"`
	// Job approver. This parameter is required if the review function is enabled.
	Approvers []*JobApprover `json:"approvers,omitempty"`
}

type Node struct {
	// Node name. The name contains a maximum of 128 characters, including only letters, numbers, hyphens (-), underscores (_), and periods (.). Names of the nodes in a job must be unique.
	Name string `json:"name" required:"true"`
	// Node type. The options are as follows:
	//
	//    Hive SQL: Runs Hive SQL scripts.
	//
	//    Spark SQL: Runs Spark SQL scripts.
	//
	//    DWS SQL: Runs DWS SQL scripts.
	//
	//    DLI SQL: Runs DLI SQL scripts.
	//
	//    Shell: Runs shell SQL scripts.
	//
	//    CDM Job: Runs CDM jobs.
	//
	//    CloudTable Manager: Manages CloudTable tables, including creating and deleting tables.
	//
	//    OBS Manager: Manages OBS paths, including creating and deleting paths.
	//
	//    RESTAPI: Sends REST API requests.
	//
	//    SMN: Sends short messages or emails.
	//
	//    MRS Spark: Runs Spark jobs of MRS.
	//
	//    MapReduce: Runs MapReduce jobs of MRS.
	//
	//    MRS Flink: Runs Flink jobs of MRS.
	//
	//    MRS HetuEngine: Runs HetuEngine jobs of MRS.
	//
	//    DLI Spark: Runs Spark jobs of DLF.
	//
	//    RDS SQL: Transfers SQL statements to RDS for execution.
	Type string `json:"type" required:"true"`
	// Location of a node on the job canvas.
	Location Location `json:"location" required:"true"`
	// Name of the previous node on which the current node depends.
	PreNodeName []string `json:"preNodeName,omitempty"`
	// Node execution condition. Whether the node is executed or not depends on the calculation result of the EL expression saved in the expression field of condition.
	Conditions []*Condition `json:"conditions,omitempty"`
	// Node properties.
	Properties []*Property `json:"properties,omitempty" required:"true"`
	// Interval at which node running results are checked.
	// Unit: second; value range: 1 to 60
	// Default value: 10
	PollingInterval int `json:"pollingInterval,omitempty"`
	// Maximum execution time of a node. If a node is not executed within the maximum execution time, the node is set to the failed state.
	// Unit: minute; value range: 5 to 1440
	// Default value: 60
	MaxExecutionTime int `json:"maxExecutionTime,omitempty"`
	// Number of the node retries. The value ranges from 0 to 5. 0 indicates no retry.
	RetryTimes int `json:"retryTimes,omitempty"`
	// Interval at which a retry is performed upon a failure. The value ranges from 5 to 120.
	RetryInterval int `json:"retryInterval,omitempty"`
	// Node failure policy.
	//    FAIL: Terminate the execution of the current job.
	//    IGNORE: Continue to execute the next node.
	//    SUSPEND: Suspend the execution of the current job.
	//    FAIL_CHILD: Terminate the execution of the subsequent node.
	//    The default value is FAIL.
	FailPolicy string `json:"failPolicy,omitempty"`
	// Event trigger for the real-time job node.
	EventTrigger *Event `json:"eventTrigger,omitempty"`
	// Cron trigger for the real-time job node.
	CronTrigger *CronTrigger `json:"cronTrigger,omitempty"`
}

type Schedule struct {
	// Scheduling type.
	//    EXECUTE_ONCE: The job runs immediately and runs only once.
	//    CRON: The job runs periodically.
	//    EVENT: The job is triggered by events.
	Type string `json:"type" required:"true"`
	// When type is set to CRON, configure the scheduling frequency and start time.
	Cron *Cron `json:"cron,omitempty"`
	// When type is set to EVENT, configure information such as the event source.
	Event *Event `json:"event,omitempty"`
}

type Param struct {
	// Name of a parameter. The name contains a maximum of 64 characters, including only letters, numbers, hyphens (-), and underscores (_).
	Name string `json:"name" required:"true"`
	// Value of the parameter. It cannot exceed 1,024 characters.
	Value string `json:"value" required:"true"`
	// Parameter type
	//    variable
	//    constants
	//    Default value: variable
	Type string `json:"type,omitempty"`
}

type Location struct {
	// Position of the node on the horizontal axis of the job canvas.
	X int `json:"x" required:"true"`
	// Position of the node on the vertical axis of the job canvas.
	Y int `json:"y" required:"true"`
}

type Condition struct {
	// Name of the previous node on which the current node depends.
	PreNodeName string `json:"preNodeName" required:"true"`
	// EL expression. If the calculation result of the EL expression is true, this node is executed.
	Expression string `json:"expression" required:"true"`
}

type CronTrigger struct {
	// Scheduling start time in the format of yyyy-MM-dd'T'HH:mm:ssZ, which is an ISO 8601 time format. For example, 2018-10-22T23:59:59+08, which indicates that a job starts to be scheduled at 23:59:59 on October 22nd, 2018.
	StartTime string `json:"startTime" required:"true"`
	// Scheduling end time in the format of yyyy-MM-dd'T'HH:mm:ssZ, which is an ISO 8601 time format. For example, 2018-10-22T23:59:59+08, which indicates that a job stops to be scheduled at 23:59:59 on October 22nd, 2018. If the end time is not set, the job will continuously be executed based on the scheduling period.
	EndTime string `json:"endTime,omitempty"`
	// Cron expression in the format of <second><minute><hour><day><month><week>.
	Expression string `json:"expression" required:"true"`
	// Time zone corresponding to the Cron expression, for example, GMT+8.
	ExpressionTimeZone string `json:"expressionTimeZone,omitempty"`
	// Job execution interval consisting of a time and time unit
	//
	// Example: 1 hours, 1 days, 1 weeks, 1 months
	Period string `json:"period" required:"true"`
	// Indicates whether to depend on the execution result of the current job's dependent job in the previous scheduling period.
	DependPrePeriod bool `json:"dependPrePeriod.,omitempty"`
	// Job dependency configuration.
	DependJobs *DependJobs `json:"dependJobs,omitempty"`
	// Number of concurrent executions allowed
	Concurrent int `json:"concurrent,omitempty"`
}

type Cron struct {
	// Scheduling start time in the format of yyyy-MM-dd'T'HH:mm:ssZ, which is an ISO 8601 time format. For example, 2018-10-22T23:59:59+08, which indicates that a job starts to be scheduled at 23:59:59 on October 22nd, 2018.
	StartTime string `json:"startTime" required:"true"`
	// Scheduling end time in the format of yyyy-MM-dd'T'HH:mm:ssZ, which is an ISO 8601 time format. For example, 2018-10-22T23:59:59+08, which indicates that a job stops to be scheduled at 23:59:59 on October 22nd, 2018. If the end time is not set, the job will continuously be executed based on the scheduling period.
	EndTime string `json:"endTime,omitempty"`
	// Cron expression in the format of <second><minute><hour><day><month><week>.
	Expression string `json:"expression" required:"true"`
	// Time zone corresponding to the Cron expression, for example, GMT+8.
	ExpressionTimeZone string `json:"expressionTimeZone,omitempty"`
	// Indicates whether to depend on the execution result of the current job's dependent job in the previous scheduling period.
	DependPrePeriod bool `json:"dependPrePeriod.,omitempty"`
	// Job dependency configuration.
	DependJobs *DependJobs `json:"dependJobs,omitempty"`
}

type Event struct {
	// Select the corresponding connection name and topic. When a new Kafka message is received, the job is triggered.
	// Set this parameter to KAFKA.
	// Event type. Currently, only newly reported data events from the DIS stream can be monitored. Each time a data record is reported, the job runs once.
	// This parameter is set to DIS.
	// Select the OBS path to be listened to. If new files exist in the path, scheduling is triggered. The path name can be referenced using variable Job.trigger.obsNewFiles. The prerequisite is that DIS notifications have been configured for the OBS path.
	// Set this parameter to OBS.
	EventType string `json:"eventType" required:"true"`
	// Job failure policy.
	//    SUSPEND: Suspend the event.
	//    IGNORE: Ignore the failure and process with the next event.
	// Default value: SUSPEND
	FailPolicy string `json:"failPolicy"`
	// Number of the concurrently scheduled jobs.
	// Value range: 1 to 128
	// Default value: 1
	Concurrent int `json:"concurrent"`
	// Access policy.
	//    LAST: Access data from the last location.
	//    NEW: Access data from a new location.
	// Default value: LAST
	ReadPolicy string `json:"readPolicy"`
}

type DependJobs struct {
	// A list of dependent jobs. Only the existing jobs can be depended on.
	Jobs []string `json:"jobs" required:"true"`
	// Dependency period.
	//    SAME_PERIOD: To run a job or not depends on the execution result of its depended job in the current scheduling period.
	//    PRE_PERIOD: To run a job or not depends on the execution result of its depended job in the previous scheduling period.
	// Default value: SAME_PERIOD
	DependPeriod string `json:"dependPeriod,omitempty"`
	// Dependency job failure policy.
	//    FAIL: Stop the job and set the job to the failed state.
	//    IGNORE: Continue to run the job.
	//    SUSPEND: Suspend the job.
	// Default value: FAIL
	DependFailPolicy string `json:"dependFailPolicy,omitempty"`
}

type Property struct {
	// Property name
	Name string `json:"name" required:"true"`
	// Property value
	Value string `json:"value" required:"true"`
}

type BasicConfig struct {
	// Job owner. The length cannot exceed 128 characters.
	Owner string `json:"owner,omitempty"`
	// Job priority. The value ranges from 0 to 2. The default value is 0. 0 indicates a top priority, 1 indicates a medium priority, and 2 indicates a low priority.
	Priority int `json:"priority,omitempty"`
	// Job execution user. The value must be an existing username.
	ExecuteUser string `json:"executeUser,omitempty"`
	// Instance timeout interval. The unit is minute. The value ranges from 5 to 1440. The default value is 60.
	InstanceTimeout int `json:"instanceTimeout,omitempty"`
	// User-defined field. The length cannot exceed 2048 characters.
	CustomFields map[string]string `json:"customFields,omitempty"`
}

type JobApprover struct {
	// Approver name.
	ApproverName string `json:"approverName" required:"true"`
}

// Create is used to create a job. A job consists of one or more nodes, such as Hive SQL and CDM Job nodes. DLF supports two types of jobs: batch jobs and real-time jobs.
// Send request POST /v1/{project_id}/jobs
func Create(client *golangsdk.ServiceClient, opts Job, workspace string) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	reqOpts := &golangsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{HeaderContentType: ApplicationJson},
	}

	if workspace != "" {
		reqOpts.MoreHeaders[HeaderWorkspace] = workspace
	}

	_, err = client.Post(client.ServiceURL(jobsEndpoint), b, nil, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
