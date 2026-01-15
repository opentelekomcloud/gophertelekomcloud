package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Dimension represents a metric dimension.
type Dimension struct {
	// Specifies the dimension name.
	Name string `json:"name" required:"true"`
	// Specifies the dimension value.
	Value string `json:"value,omitempty"`
}

// AddOpts contains the options for adding resources to an alarm rule.
type AddOpts struct {
	// Specifies the list of resources to add.
	// A maximum of 1000 resources can be added at a time.
	Resources [][]Dimension `json:"resources" required:"true"`
}

// Add adds resources to an alarm rule in batches.
func Add(client *golangsdk.ServiceClient, alarmId string, opts AddOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/alarms/{alarm_id}/resources/batch-create
	_, err = client.Post(client.ServiceURL("alarms", alarmId, "resources", "batch-create"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
