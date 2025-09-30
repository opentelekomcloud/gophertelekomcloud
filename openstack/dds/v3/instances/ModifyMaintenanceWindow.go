package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ModifyMaintenanceWindowOpt struct {
	StartTime  string `json:"start_time" required:"true"`
	EndTime    string `json:"end_time" required:"true"`
	InstanceId string `json:"-"`
}

func ModifyMaintenanceWindow(client *golangsdk.ServiceClient, opts ModifyMaintenanceWindowOpt) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "maintenance-window"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
