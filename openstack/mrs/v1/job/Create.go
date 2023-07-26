package job

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
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
	ClusterId string `json:"cluster_id" required:"true"`
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
func Create(c *golangsdk.ServiceClient, opts any) (*JobExecution, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.1/{project_id}/jobs/submit-job
	raw, err := c.Post(c.ServiceURL("jobs", "submit-job"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})

	return extra(err, raw)
}
