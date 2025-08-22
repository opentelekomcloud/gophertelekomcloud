package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func UpdatePass(client *golangsdk.ServiceClient, instanceId string, password string) error {
	opts := struct {
		Password string `json:"password"`
	}{Password: password}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("instances", instanceId, "password"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}
