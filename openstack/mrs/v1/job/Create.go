package job

import (
	"log"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type CreateOpts struct {
	JobType        int    `json:"job_type" required:"true"`
	JobName        string `json:"job_name" required:"true"`
	ClusterID      string `json:"cluster_id" required:"true"`
	JarPath        string `json:"jar_path" required:"true"`
	Arguments      string `json:"arguments,omitempty"`
	Input          string `json:"input,omitempty"`
	Output         string `json:"output,omitempty"`
	JobLog         string `json:"job_log,omitempty"`
	HiveScriptPath string `json:"hive_script_path,omitempty"`
	IsProtected    bool   `json:"is_protected,omitempty"`
	IsPublic       bool   `json:"is_public,omitempty"`
}

type CreateOptsBuilder interface {
	ToJobCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToJobCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToJobCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	log.Printf("[DEBUG] create url:%q, body=%#v", c.ServiceURL("jobs/submit-job"), b)
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200},
		MoreHeaders: RequestOpts.MoreHeaders}
	_, r.Err = c.Post(c.ServiceURL("jobs/submit-job"), b, &r.Body, reqOpt)
	return
}
