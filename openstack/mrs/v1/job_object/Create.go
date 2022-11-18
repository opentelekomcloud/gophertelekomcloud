package job_object

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Job type code
	// 1: MapReduce
	// 2: Spark
	// 3: Hive Script
	// 4: HiveQL (not supported currently)
	// 5: DistCp, importing and exporting data. For details, see CreateDistCpOpts.
	// 6: Spark Script
	// 7: Spark SQL, submitting Spark SQL statements. For details, see CreateSparkOpts. (Not supported in this API currently.)
	// NOTE:
	// Spark and Hive jobs can be added to only clusters that include Spark and Hive components.
	JobType int `json:"job_type" required:"true"`
	// Job name
	// Contains only 1 to 64 letters, digits, hyphens (-), and underscores (_).
	// NOTE:
	// Identical job names are allowed but not recommended.
	JobName string `json:"job_name" required:"true"`
	// Cluster ID
	ClusterID string `json:"cluster_id" required:"true"`
	// Path of the JAR or SQL file for program execution
	// The parameter must meet the following requirements:
	// Contains a maximum of 1,023 characters, excluding special characters such as ;|&><'$.
	// The address cannot be empty or full of spaces.
	// Starts with / or s3a://. The OBS path does not support files or programs encrypted by KMS.
	// Spark Script must end with .sql while MapReduce and Spark Jar must end with .jar.sql and jar are case-insensitive.
	JarPath string `json:"jar_path" required:"true"`
	// Key parameter for program execution. The parameter is specified by the function of the user's program.
	// MRS is only responsible for loading the parameter.
	// The parameter contains a maximum of 2,047 characters, excluding special characters such as ;|&>'<$, and can be left blank.
	// NOTE:
	// When entering a parameter containing sensitive information (for example, login password),
	// you can add an at sign (@) before the parameter name to encrypt the parameter value.
	// This prevents the sensitive information from being persisted in plaintext.
	// Therefore, when you view job information on the MRS, sensitive information will be displayed as asterisks (*).
	// For example, username=admin @password=admin_123.
	Arguments string `json:"arguments,omitempty"`
	// Path for inputting data, which must start with / or s3a://. Set this parameter to a correct OBS path.
	// The OBS path does not support files or programs encrypted by KMS.
	// The parameter contains a maximum of 1,023 characters, excluding special characters such as ;|&>'<$, and can be left blank.
	Input string `json:"input,omitempty"`
	// Path for outputting data, which must start with / or s3a://. A correct OBS path is required.
	// If the path does not exist, the system automatically creates it.
	// The parameter contains a maximum of 1,023 characters, excluding special characters such as ;|&>'<$, and can be left blank.
	Output string `json:"output,omitempty"`
	// Path for storing job logs that record job running status. The path must start with / or s3a://. A correct OBS path is required.
	// The parameter contains a maximum of 1,023 characters, excluding special characters such as ;|&>'<$, and can be left blank.
	JobLog string `json:"job_log,omitempty"`
	// SQL program path
	// This parameter is needed by Spark Script and Hive Script jobs only, and must meet the following requirements:
	// Contains a maximum of 1,023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces.
	// The path must start with / or s3a://. The OBS path does not support files or programs encrypted by KMS.
	// The path must end with .sql.sql is case-insensitive.
	HiveScriptPath string `json:"hive_script_path,omitempty"`

	IsProtected bool `json:"is_protected,omitempty"`
	IsPublic    bool `json:"is_public,omitempty"`
}

type CreateDistCpOpts struct {
	// 5: DistCp, importing and exporting data.
	JobType int `json:"job_type"`
	// Job name
	// Contains only 1 to 64 letters, digits, hyphens (-), and underscores (_).
	// NOTE:
	// Identical job names are allowed but not recommended.
	JobName   string `json:"job_name"`
	ClusterId string `json:"cluster_id"`
	// Data source path
	// When you import data, the parameter is set to an OBS path. Files or programs encrypted by KMS are not supported.
	// When you export data, the parameter is set to an HDFS path.
	Input string `json:"input"`
	// Data receiving path
	// When you import data, the parameter is set to an HDFS path.
	// When you export data, the parameter is set to an OBS path.
	Output string `json:"output"`
	// Types of file operations, including:
	// export: Export data from HDFS to OBS.
	// import: Import data from OBS to HDFS.
	FileAction string `json:"file_action"`
}

type CreateSparkOpts struct {
	JobType    int    `json:"job_type"`
	JobName    string `json:"job_name"`
	ClusterId  string `json:"cluster_id"`
	JarPath    string `json:"jar_path"`
	Arguments  string `json:"arguments"`
	Input      string `json:"input"`
	Output     string `json:"output"`
	JobLog     string `json:"job_log"`
	FileAction string `json:"file_action"`
	// Spark SQL statement, which needs Base64 encoding and decoding. ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/
	// is a standard encoding table. MRS uses ABCDEFGHILKJMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/ for Base64 encoding.
	// The value of the hql parameter is generated by adding any letter to the beginning of the encoded character string.
	// The Spark SQL statement is generated by decoding the value in the background.
	// Example:
	// Obtain the Base64 encoding tool.
	// Enter the show tables; Spark SQL statement in the encoding tool to perform Base64 encoding.
	// Obtain the encoded character string c2hvdyB0YWLsZXM7.
	// At the beginning of c2hvdyB0YWLsZXM7, add any letter, for example, g. Then, the character string becomes
	// gc2hvdyB0YWLsZXM7, that is, the value of the hql parameter.
	Hql            string `json:"hql"`
	HiveScriptPath string `json:"hive_script_path"`
}

// Create Use CreateOpts or CreateDistCpOpts or CreateSparkOpts
func Create(c *golangsdk.ServiceClient, opts interface{}) (*JobExecution, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.1/{project_id}/jobs/submit-job
	raw, err := c.Post(c.ServiceURL("jobs", "submit-job"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})
	if err != nil {
		return nil, err
	}

	var res JobExecution
	err = extract.IntoStructPtr(raw.Body, &res, "job_execution")
	return &res, err
}

type JobExecution struct {
	// Whether job execution objects are generated by job templates.
	Templated bool `json:"templated,omitempty"`
	// Creation time, which is a 10-bit timestamp.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Update time, which is a 10-bit timestamp.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Job ID
	Id string `json:"id,omitempty"`
	// Project ID. For details on how to obtain the project ID
	TenantId string `json:"tenant_id,omitempty"`
	// Job application ID
	JobId string `json:"job_id,omitempty"`
	// Job name
	JobName string `json:"job_name,omitempty"`
	// Data input ID
	InputId string `json:"input_id,omitempty"`
	// Data output ID
	OutputId string `json:"output_id,omitempty"`
	// Start time of job execution, which is a 10-bit timestamp.
	StartTime int64 `json:"start_time,omitempty"`
	// End time of job execution, which is a 10-bit timestamp.
	EndTime int64 `json:"end_time,omitempty"`
	// Cluster ID
	ClusterId string `json:"cluster_id,omitempty"`
	// Workflow ID of Oozie
	EngineJobId string `json:"engine_job_id,omitempty"`
	// Returned code for an execution result
	ReturnCode string `json:"return_code,omitempty"`
	// Whether a job is public
	// The current version does not support this function.
	IsPublic bool `json:"is_public,omitempty"`
	// Whether a job is protected
	// The current version does not support this function.
	IsProtected bool `json:"is_protected,omitempty"`
	// Group ID of a job
	GroupId string `json:"group_id,omitempty"`
	// Path of the .jar file for program execution
	JarPath string `json:"jar_path,omitempty"`
	// Address for inputting data
	Input string `json:"input,omitempty"`
	// Address for outputting data
	Output string `json:"output,omitempty"`
	// Address for storing job logs
	JobLog string `json:"job_log,omitempty"`
	// Job type code
	// 1: MapReduce
	// 2: Spark
	// 3: Hive Script
	// 4: HiveQL (not supported currently)
	// 5: DistCp
	// 6: Spark Script
	// 7: Spark SQL (not supported in this API currently)
	JobType int32 `json:"job_type,omitempty"`
	// Data import and export
	FileAction string `json:"file_action,omitempty"`
	// Key parameter for program execution. The parameter is specified by the function of the user's internal program.
	// MRS is only responsible for loading the parameter. This parameter can be empty.
	Arguments string `json:"arguments,omitempty"`
	// Job status code
	// -1: Terminated
	// 1: Starting
	// 2: Running
	// 3: Completed
	// 4: Abnormal
	// 5: Error
	JobState int32 `json:"job_state,omitempty"`
	// Final job status
	// 0: unfinished
	// 1: terminated due to an execution error
	// 2: executed successfully
	// 3: canceled
	JobFinalStatus int32 `json:"job_final_status,omitempty"`
	// Address of the Hive script
	HiveScriptPath string `json:"hive_script_path,omitempty"`
	// User ID for creating jobs
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	CreateBy string `json:"create_by,omitempty"`
	// Number of completed steps
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	FinishedStep int32 `json:"finished_step,omitempty"`
	// Main ID of a job
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	JobMainId string `json:"job_main_id,omitempty"`
	// Step ID of a job
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	JobStepId string `json:"job_step_id,omitempty"`
	// Delay time, which is a 10-bit timestamp.
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	PostponeAt int64 `json:"postpone_at,omitempty"`
	// Step name of a job
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	StepName string `json:"step_name,omitempty"`
	// Number of steps
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	StepNum int32 `json:"step_num,omitempty"`
	// Number of tasks
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	TaskNum int32 `json:"task_num,omitempty"`
	// User ID for updating jobs
	UpdateBy string `json:"update_by,omitempty"`
	// Token
	// The current version does not support this function.
	Credentials string `json:"credentials,omitempty"`
	// User ID for creating jobs
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	UserId string `json:"user_id,omitempty"`
	// Key-value pair set for saving job running configurations
	JobConfigs map[string]interface{} `json:"job_configs,omitempty"`
	// Authentication information
	// The current version does not support this function.
	Extra map[string]interface{} `json:"extra,omitempty"`
	// Data source URL
	DataSourceUrls map[string]interface{} `json:"data_source_urls,omitempty"`
	// Key-value pair set, containing job running information returned by Oozie
	Info map[string]interface{} `json:"info,omitempty"`
}
