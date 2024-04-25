package job

import (
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const statusEndpoint = "status"

// GetRunningStatus is used to view running status of a real-time job.
// Send request GET /v1/{project_id}/jobs/{job_name}/status
func GetRunningStatus(client *golangsdk.ServiceClient, jobName, workspace string) (*RunningStatusResp, error) {

	var opts *golangsdk.RequestOpts
	if workspace != "" {
		opts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	raw, err := client.Get(client.ServiceURL(jobsEndpoint, jobName, statusEndpoint), nil, opts)
	if err != nil {
		return nil, err
	}

	var res *RunningStatusResp
	err = extract.Into(raw.Body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type RunningStatusResp struct {
	// Name of a solution.
	Name string `json:"name"`
	// Node status list.
	Nodes []*NodeStatusResp `json:"nodes,omitempty"`
	// Job status.
	//    STARTING
	//    NORMAL
	//    EXCEPTION
	//    STOPPING
	//    STOPPED
	Status string `json:"status,omitempty"`
	// Start time.
	StartTime time.Time `json:"startTime"`
	// End time.
	EndTime *time.Time `json:"endTime,omitempty"`
	// Last update time.
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
}

type NodeStatusResp struct {
	// Node name.
	Name string `json:"name"`
	// Node status.
	//    STARTING
	//    NORMAL
	//    EXCEPTION
	//    STOPPING
	//    STOPPED
	Status string `json:"status,omitempty"`
	// Path for storing node run logs.
	LogPath string `json:"logPath,omitempty"`
	// Node type.
	//    Hive SQL: Runs Hive SQL scripts.
	//    Spark SQL: Runs Spark SQL scripts.
	//    DWS SQL: Runs DWS SQL scripts.
	//    DLI SQL: Runs DLI SQL scripts.
	//    Shell: Runs shell SQL scripts.
	//    CDM Job: Runs CDM jobs.
	//    DIS Transfer Task: Creates DIS dump tasks.
	//    CS Job: Creates and starts CloudStream jobs.
	//    CloudTable Manager: Manages CloudTable tables, including creating and deleting tables.
	//    OBS Manager: Manages OBS paths, including creating and deleting paths.
	//    RESTAPI: Sends REST API requests.
	//    SMN: Sends short messages or emails.
	//    MRS Spark: Runs Spark jobs of MRS.
	//    MapReduce: Runs MapReduce jobs of MRS.
	Type string `json:"type"`
}
