package disk

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Detach(client *golangsdk.ServiceClient, serverID, volumeID string, deleteFlag int) (*JobResponse, error) {
	url := client.ServiceURL("cloudservers", serverID, "detachvolume", volumeID)
	if deleteFlag != 0 {
		url += "?delete_flag=1"
	}

	raw, err := client.Delete(url, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res JobResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
