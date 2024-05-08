package connection

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// IsCheck indicates whether to perform check. The default value is No.
	IsCheck bool `q:"ischeck"`
}

// Update is used to edit a connection.
// Send request PUT /v1/{project_id}/connections/{connection_name}?ischeck=true
func Update(client *golangsdk.ServiceClient, conn Connection, opts UpdateOpts) error {

	url, err := golangsdk.NewURLBuilder().WithEndpoints(connectionsEndpoint, conn.Name).WithQueryParams(&opts).Build()
	if err != nil {
		return err
	}

	b, err := build.RequestBody(conn, "")
	if err != nil {
		return err
	}

	reqOpts := &golangsdk.RequestOpts{
		OkCodes: []int{204},
	}

	if conn.Workspace != "" {
		reqOpts.MoreHeaders = map[string]string{HeaderWorkspace: conn.Workspace}
	}

	_, err = client.Put(url.String(), b, nil, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
