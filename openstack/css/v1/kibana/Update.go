package kibana

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func Update(client *golangsdk.ServiceClient, clusterId string, size int) error {
	request := UpdateOpts{
		Bandwidth: Bandwidth{
			Size: size,
		},
	}

	b, err := build.RequestBody(request, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterId, "publickibana", "bandwidth"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}

type UpdateOpts struct {
	Bandwidth Bandwidth `json:"bandWidth" required:"true"`
}

type Bandwidth struct {
	Size int `json:"size" required:"true"`
}
