package evacuate

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// EvacuateResult is the response from an Evacuate operation.
// Call its ExtractAdminPass method to retrieve the admin password of the instance.
// The admin password will be an empty string if the cloud is not configured to inject admin passwords..
type EvacuateResult struct {
	golangsdk.Result
}

func (raw EvacuateResult) ExtractAdminPass() (string, error) {
	var res struct {
		AdminPass string `json:"adminPass"`
	}
	err = extract.Into(raw.Body, &res)
	if err != nil && err.Error() == "EOF" {
		return "", nil
	}
	return &res, err
}
