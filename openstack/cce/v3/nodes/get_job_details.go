package nodes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetJobDetails retrieves a particular job based on its unique ID
func GetJobDetails(client *golangsdk.ServiceClient, jobID string) (*Job, error) {
	raw, err := client.Get(client.ServiceURL("jobs", jobID), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts, JSONBody: nil,
	})
	if err != nil {
		return nil, err
	}

	var res Job
	err = extract.Into(raw, &res)
	return &res, err
}

// Job Structure
type Job struct {
	// API type, fixed value "Job"
	Kind string `json:"kind"`
	// API version, fixed value "v3"
	Apiversion string `json:"apiVersion"`
	// Node metadata
	Metadata JobMetadata `json:"metadata"`
	// Node detailed parameters
	Spec JobSpec `json:"spec"`
	// Node status information
	Status JobStatus `json:"status"`
}

type JobMetadata struct {
	// ID of the job
	ID string `json:"uid"`
}

type JobSpec struct {
	// Type of job
	Type string `json:"type"`
	// ID of the cluster where the job is located
	ClusterID string `json:"clusterUID"`
	// ID of the IaaS resource for the job operation
	ResourceID string `json:"resourceID"`
	// The name of the IaaS resource for the job operation
	ResourceName string `json:"resourceName"`
	// List of child jobs
	SubJobs []Job `json:"subJobs"`
	// ID of the parent job
	OwnerJob string `json:"ownerJob"`
}

type JobStatus struct {
	// Job status
	Phase string `json:"phase"`
	// The reason why the job becomes the current state
	Reason string `json:"reason"`
	// The job becomes the current state details
	Message string `json:"message"`
}
