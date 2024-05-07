package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const (
	NormalJob   = "NORMAL_JOB"
	BatchJob    = "BATCH_JOB"
	ScenarioJob = "SCENARIO_JOB"
)

type CreateRandomOpts struct {
	// Jobs is a jobs list.
	Jobs []Job `json:"jobs" required:"true"`
	// Clusters is a list of IDs of CDM clusters. The system selects a random cluster in running state from the specified clusters and creates and executes a migration job in the cluster.
	Clusters []string `json:"clusters" required:"true"`
}

// CreateRandom is used to create and execute a job in a random cluster.
// Send request POST /v1.1/{project_id}/clusters/job
func CreateRandom(client *golangsdk.ServiceClient, opts CreateRandomOpts, xLang string) (*JobResp, error) {

	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			HeaderContentType: ApplicationJson,
			HeaderXLanguage:   xLang,
		},
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(clustersEndpoint, jobEndpoint), b, nil, reqOpts)

	if err != nil {
		return nil, err
	}

	var resp JobResp
	err = extract.Into(raw.Body, &resp)
	return &resp, err
}

// Job represents a job configuration
type Job struct {
	// JobType specifies the type of job (e.g., NORMAL_JOB, BATCH_JOB, SCENARIO_JOB).
	JobType string `json:"job_type,omitempty"`

	// FromConnectorName is the source link type (required).
	FromConnectorName string `json:"from-connector-name" required:"true"`

	// ToConfigValues holds destination link parameter configuration.
	ToConfigValues *ConfigValues `json:"to-config-values"  required:"true"`

	// ToLinkName is the destination link name (required).
	ToLinkName string `json:"to-link-name" required:"true"`

	// DriverConfigValues holds job parameter configuration.
	DriverConfigValues *ConfigValues `json:"driver-config-values" required:"true"`

	// FromConfigValues holds source link parameter configuration.
	FromConfigValues *ConfigValues `json:"from-config-values" required:"true"`

	// ToConnectorName is the destination link type (optional).
	ToConnectorName string `json:"to-connector-name,omitempty"`

	// Name is the job name with a limit of 1-240 characters (optional).
	Name string `json:"name,omitempty"`

	// FromLinkName is the source link name (optional).
	FromLinkName string `json:"from-link-name,omitempty"`

	// CreationUser is the user who created the job (optional).
	CreationUser string `json:"creation-user,omitempty"`

	// CreationDate is the time when the job was created (optional).
	CreationDate int64 `json:"creation-date,omitempty"`

	// UpdateDate is the time when the job was last updated (optional).
	UpdateDate int64 `json:"update-date,omitempty"`

	// IsIncremental indicates whether the job is incremental (optional).
	IsIncremental bool `json:"is_incre_job,omitempty"`

	// Flag is a job flag (optional).
	Flag int `json:"flag,omitempty"`

	// FilesRead is the number of files read during the job (optional).
	FilesRead int `json:"files_read,omitempty"`

	// UpdateUser is the user who last updated the job (optional).
	UpdateUser string `json:"update-user,omitempty"`

	// ExternalID is an external ID associated with the job (optional).
	ExternalID string `json:"external_id,omitempty"`

	// Type is the task type (optional).
	Type string `json:"type,omitempty"`

	// ExecuteStartDate is the execution start date (optional).
	ExecuteStartDate int64 `json:"execute_start_date,omitempty"`

	// DeleteRows is the number of rows deleted during the job (optional).
	DeleteRows int `json:"delete_rows,omitempty"`

	// Enabled indicates whether the link is activated (optional).
	Enabled bool `json:"enabled,omitempty"`

	// BytesWritten is the number of bytes written during the job (optional).
	BytesWritten int64 `json:"bytes_written,omitempty"`

	// ID is the job ID (optional).
	ID int `json:"id,omitempty"`

	// IsUseSQL indicates whether to use SQL statements (optional).
	IsUseSQL bool `json:"is_use_sql,omitempty"`

	// UpdateRows is the number of rows updated during the job (optional).
	UpdateRows int `json:"update_rows,omitempty"`

	// GroupName is the group name associated with the job (optional).
	GroupName string `json:"group_name,omitempty"`

	// BytesRead is the number of bytes read during the job (optional).
	BytesRead int64 `json:"bytes_read,omitempty"`

	// ExecuteUpdateDate is the execution update date (optional).
	ExecuteUpdateDate int64 `json:"execute_update_date,omitempty"`

	// WriteRows is the number of rows written during the job (optional).
	WriteRows int `json:"write_rows,omitempty"`

	// RowsWritten is the number of rows written during the job (optional).
	RowsWritten int `json:"rows_written,omitempty"`

	// RowsRead is the number of rows read during the job (optional).
	RowsRead int64 `json:"rows_read,omitempty"`

	// FilesWritten is the number of written files (optional).
	FilesWritten int64 `json:"files_written,omitempty"`

	// IsIncrementing indicates incremental or not (optional).
	IsIncrementing bool `json:"is_incrementing,omitempty"`

	// ExecuteCreateDate is the execution creation date (optional).
	ExecuteCreateDate int64 `json:"execute_create_date,omitempty"`

	// Status is Job execution status. Can be one of: BOOTING, RUNNING, SUCCEEDED, FAILED, NEW. Optional.
	Status int64 `json:"status,omitempty"`
}

type ConfigValues struct {
	// Configs The data structures of source link parameters, destination link parameters, and job parameters are the same. However, the inputs parameter varies.
	Configs []*Config `json:"configs" required:"true"`

	// ExtendedConfigs is extended configuration.
	ExtendedConfigs *ExtendedConfigs `json:"extended-configs,omitempty"`
}

type Config struct {
	// Inputs is an Input parameter list. Each element in the list is in name,value format. For details, see the descriptions of inputs parameters.
	// In the from-config-values data structure, the value of this parameter varies with the source link type.
	// For details, see section "Source Job Parameters" in the Cloud Data Migration User Guide. In the to-cofig-values data structure, the value of this parameter varies with the destination link type.
	// For details, see section "Destination Job Parameters" in the Cloud Data Migration User Guide.
	// For details about the inputs parameter in the driver-config-values data structure, see the job parameter descriptions.
	Inputs []*Input `json:"inputs" required:"true"`

	// Name is a configuration name. The value is fromJobConfig for a source job, toJobConfig for a destination job, and linkConfig for a link.
	Name string `json:"name"  required:"true"`

	// Id is a configuration ID.
	Id int `json:"id,omitempty"`

	// Type is a configuration type.
	Type string `json:"type,omitempty"`
}

type Input struct {
	Name  string `json:"name"  required:"true"`
	Value string `json:"value" required:"true"`
	Type  string `json:"type,omitempty"`
}

type ExtendedConfigs struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Response

// Response holds job running information
type JobResp struct {
	// Submissions is an array of StartJobSubmission objects.
	Submissions []*StartJobSubmission `json:"submissions"`
}

// StartJobSubmission represents a job submission
type StartJobSubmission struct {
	// IsIncrementing indicates whether the job migrates incremental data.
	IsIncrementing bool `json:"isIncrementing"`

	// DeleteRows is the number of rows deleted during the job.
	DeleteRows int `json:"delete_rows"`

	// UpdateRows is the number of rows updated during the job.
	UpdateRows int `json:"update_rows"`

	// WriteRows is the number of rows written during the job.
	WriteRows int `json:"write_rows"`

	// SubmissionID is the ID of the submitted job.
	SubmissionID int `json:"submission-id"`

	// JobName is the name of the job.
	JobName string `json:"job-name"`

	// CreationUser is the user who created the job.
	CreationUser string `json:"creation-user"`

	// CreationDate is the time when the job was created, accurate to millisecond.
	CreationDate int64 `json:"creation-date"`

	// ExecuteDate is the time when the job was executed.
	ExecuteDate int64 `json:"execute-date"`

	// Progress is the job progress. If a job fails, the value is -1. Otherwise, the value ranges from 0 to 100.
	Progress float64 `json:"progress"`

	// Status represents the job's execution status.
	Status string `json:"status"`

	// IsStopingIncrement indicates whether to stop incremental data migration.
	IsStopingIncrement string `json:"isStopingIncrement"`

	// IsExecuteAuto indicates whether to execute the job as scheduled.
	IsExecuteAuto bool `json:"is-execute-auto"`

	// LastUpdateDate is the time when the job was last updated.
	LastUpdateDate int64 `json:"last-update-date"`

	// LastUpdateUser is the user who last updated the job status.
	LastUpdateUser string `json:"last-udpate-user"`

	// IsDeleteJob indicates whether to delete the job after it is executed.
	IsDeleteJob bool `json:"isDeleteJob"`
}
