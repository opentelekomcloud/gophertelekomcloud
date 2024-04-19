package script

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteReq struct {
	Approvers []*JobApprover `json:"approvers,omitempty"`
}

// Delete is used to delete a specific script.
// Send request DELETE /v1/{project_id}/scripts/{script_name}
func Delete(client *golangsdk.ServiceClient, scriptName, workspace string, opts *DeleteReq) error {
	var err error
	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{},
		OkCodes:     []int{204},
	}

	var b *build.Body
	if opts != nil {
		b, err = build.RequestBody(opts, "")
		if err != nil {
			return err
		}

		reqOpts.MoreHeaders[HeaderContentType] = ApplicationJson
	}

	if workspace != "" {
		reqOpts.MoreHeaders[HeaderWorkspace] = workspace
	}

	_, err = client.DeleteWithBody(client.ServiceURL(scriptsURL, scriptName), b, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
