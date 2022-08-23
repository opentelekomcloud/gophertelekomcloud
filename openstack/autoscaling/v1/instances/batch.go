package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type BatchOpts struct {
	Instances   []string `json:"instances_id" required:"true"`
	IsDeleteEcs string   `json:"instance_delete,omitempty"`
	Action      string   `json:"action,omitempty"`
}

func batch(client *golangsdk.ServiceClient, groupID string, opts BatchOpts) error {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("scaling_group_instance", groupID, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}

func BatchAdd(client *golangsdk.ServiceClient, groupID string, instances []string) error {
	return batch(client, groupID, BatchOpts{
		Instances: instances,
		Action:    "ADD",
	})
}

func BatchDelete(client *golangsdk.ServiceClient, groupID string, instances []string, deleteEcs string) error {
	return batch(client, groupID, BatchOpts{
		Instances:   instances,
		IsDeleteEcs: deleteEcs,
		Action:      "REMOVE",
	})
}
