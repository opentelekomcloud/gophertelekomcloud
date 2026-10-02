package diagnosis

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Stack contains the stack information of a client.
type Stack struct {
	// Thread name.
	ThreadName string `json:"thread_name"`
	// Stack information of the client.
	Stack string `json:"stack"`
}

// GetStack queries stack information.
// The stack ID is obtained from NodeReport.StackID of a diagnosis report.
// Send GET to /v2/{project_id}/rocketmq/diagnosis/stack/{stack_id}
func GetStack(client *golangsdk.ServiceClient, stackID string) (*Stack, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("rocketmq", "diagnosis", "stack", stackID).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Stack
	err = extract.Into(raw.Body, &res)
	return &res, err
}
