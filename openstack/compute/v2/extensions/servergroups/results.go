package servergroups

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// A ServerGroup creates a policy for instance placement in the cloud.
type ServerGroup struct {
	// ID is the unique ID of the Server Group.
	ID string `json:"id"`
	// Name is the common name of the server group.
	Name string `json:"name"`
	// Polices are the group policies.
	// Normally a single policy is applied:
	// "affinity" will place all servers within the server group on the same compute node.
	// "anti-affinity" will place servers within the server group on different compute nodes.
	Policies []string `json:"policies"`
	// Members are the members of the server group.
	Members []string `json:"members"`
	// Metadata includes a list of all user-specified key-value pairs attached to the Server Group.
	Metadata map[string]interface{}
}

func extra(err error, raw *http.Response) (*ServerGroup, error) {
	if err != nil {
		return nil, err
	}

	var res ServerGroup
	err = extract.IntoStructPtr(raw.Body, &res, "server_group")
	return &res, err
}
