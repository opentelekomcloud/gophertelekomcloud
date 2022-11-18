package job_object

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(c *golangsdk.ServiceClient, id string) (*Job, error) {
	// GET /v1.1/{project_id}/job-exes/{job_exe_id}
	raw, err := c.Get(c.ServiceURL("job-exes", id), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})
	if err != nil {
		return nil, err
	}

	var res Job
	err = extract.IntoStructPtr(raw.Body, &res, "job_execution")
	return &res, err
}

type Job struct {
	ID             string `json:"id"`
	TenantID       string `json:"tenant_id"`
	JobID          string `json:"job_id"`
	JobName        string `json:"job_name"`
	StartTime      int    `json:"start_time"`
	EndTime        int    `json:"end_time"`
	ClusterID      string `json:"cluster_id"`
	GroupID        string `json:"group_id"`
	JarPath        string `json:"jar_path"`
	Input          string `json:"input"`
	Output         string `json:"output"`
	JobLog         string `json:"job_log"`
	JobType        int    `json:"job_type"`
	FileAction     string `json:"file_action"`
	Arguments      string `json:"arguments"`
	Hql            string `json:"hql"`
	JobState       int    `json:"job_state"`
	JobFinalStatus int    `json:"job_final_status"`
	HiveScriptPath string `json:"hive_script_path"`
	CreateBy       string `json:"create_by"`
	FinishedStep   int    `json:"finished_step"`
	JobMainID      string `json:"job_main_id"`
	JobStepID      string `json:"job_step_id"`
	PostponeAt     int    `json:"postpone_at"`
	StepName       string `json:"step_name"`
	StepNum        int    `json:"step_num"`
	TaskNum        int    `json:"task_num"`
	UpdateBy       string `json:"update_by"`
	SpendTime      int    `json:"spend_time"`
	StepSeq        int    `json:"step_seq"`
	Progress       string `json:"progress"`
}
