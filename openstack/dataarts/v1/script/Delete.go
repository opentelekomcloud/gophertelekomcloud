package script

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteReq struct {
	// Workspace ID.
	Workspace string `json:"-"`
	// Approvers is a script approver. This parameter is required if the review function is enabled.
	Approvers []*JobApprover `json:"approvers,omitempty"`
}

// Delete is used to delete a specific script.
// Send request DELETE /v1/{project_id}/scripts/{script_name}
func Delete(client *golangsdk.ServiceClient, scriptName string, opts DeleteReq) error {
	if opts.Workspace == "" {
		return fmt.Errorf("workspace is required")
	}
	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{HeaderWorkspace: opts.Workspace},
		OkCodes:     []int{204},
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	reqOpts.MoreHeaders[HeaderContentType] = ApplicationJson

	_, err = client.DeleteWithBody(client.ServiceURL(scriptsEndpoint, scriptName), b, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
