package attachinterfaces

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List makes a request against the nova API to list the server's interfaces.
func List(client *golangsdk.ServiceClient, serverID string) ([]Interface, error) {
	raw, err := client.Get(client.ServiceURL("servers", serverID, "os-interface"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Interface
	err = extract.IntoSlicePtr(raw.Body, &res, "interfaceAttachments")
	return res, err
}

// FixedIP represents a Fixed IP Address.
// This struct is also used when creating an attachment,
// but it is not possible to specify a SubnetID.
type FixedIP struct {
	SubnetID  string `json:"subnet_id,omitempty"`
	IPAddress string `json:"ip_address"`
}
