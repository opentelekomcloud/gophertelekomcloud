package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type BatchOptsBuilder interface {
	ToInstanceBatchMap() (map[string]interface{}, error)
}

type BatchOpts struct {
	Instances   []string `json:"instances_id" required:"true"`
	IsDeleteEcs string   `json:"instance_delete,omitempty"`
	Action      string   `json:"action,omitempty"`
}

func (opts BatchOpts) ToInstanceBatchMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func batch(client *golangsdk.ServiceClient, groupID string, opts BatchOptsBuilder) (r BatchResult) {
	b, err := opts.ToInstanceBatchMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("scaling_group_instance", groupID, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

func BatchAdd(client *golangsdk.ServiceClient, groupID string, instances []string) (r BatchResult) {
	var opts = BatchOpts{
		Instances: instances,
		Action:    "ADD",
	}
	return batch(client, groupID, opts)
}

func BatchDelete(client *golangsdk.ServiceClient, groupID string, instances []string, deleteEcs string) (r BatchResult) {
	var opts = BatchOpts{
		Instances:   instances,
		IsDeleteEcs: deleteEcs,
		Action:      "REMOVE",
	}
	return batch(client, groupID, opts)
}
